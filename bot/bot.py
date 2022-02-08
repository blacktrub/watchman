import uuid

import aiosqlite
from aiogram import Bot, executor, Dispatcher, types


def get_token_from_file() -> str:
    with open(".token") as f:
        return f.read().strip()


API_TOKEN = get_token_from_file()

DB_NAME = "storage.db"

bot = Bot(token=API_TOKEN)
dp = Dispatcher(bot)


COMMANDS = {
    "/start": "Start bot",
    "/help": "Show this message",
    "/add": "Add new project",
    "/remove": "Remove project",
}


def make_project_hash() -> str:
    return uuid.uuid4().hex


@dp.message_handler(commands=["start"])
async def hello(msg: types.Message):
    await msg.answer("Hello, chat!")


@dp.message_handler(commands=["help"])
async def help(msg: types.Message):
    await msg.answer("\n".join((f"{k} - {v}" for k, v in COMMANDS.items())))


@dp.message_handler(commands=["add"])
async def add_project(msg: types.Message):
    telegram_id = msg.from_user.id
    async with aiosqlite.connect(f"../{DB_NAME}") as db:
        # TODO: check if user exists
        cur = await db.execute(
            "insert into user (telegram_id) values (?)", [telegram_id]
        )
        user_id = cur.lastrowid
        # TODO: check if projects exists
        await db.execute(
            "insert into project (hash, user_id) values (?, ?)",
            [
                user_id,
                make_project_hash(),
            ],
        )
        await db.commit()

    # TODO: return help text how to use project
    # TODO: return an url for project
    await msg.answer("Project was created")


if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
