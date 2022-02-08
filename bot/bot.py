from aiogram import Bot, executor, Dispatcher, types


def get_token_from_file() -> str:
    with open(".token") as f:
        return f.read().strip()


API_TOKEN = get_token_from_file()

bot = Bot(token=API_TOKEN)
dp = Dispatcher(bot)


@dp.message_handler(commands=["start", "help"])
async def hello(msg: types.Message):
    await msg.reply("Hello, chat!")


if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)

