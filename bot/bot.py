from aiogram import Bot, executor, Dispatcher, types


def get_token_from_file() -> str:
    with open(".token") as f:
        return f.read().strip()


API_TOKEN = get_token_from_file()

bot = Bot(token=API_TOKEN)
dp = Dispatcher(bot)


COMMANDS = {
    "/start": "Start bot",
    "/help": "Show this message",
    "/add": "Add new project",
    "/remove": "Remove project",
}


@dp.message_handler(commands=["start"])
async def hello(msg: types.Message):
    await msg.answer("Hello, chat!")


@dp.message_handler(commands=["help"])
async def help(msg: types.Message):
    await msg.answer("\n".join((f"{k} - {v}" for k, v in COMMANDS.items())))


if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
