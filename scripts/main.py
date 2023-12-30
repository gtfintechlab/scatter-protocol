from pymongo import MongoClient
import os


def db_connect():
    if MongoClient().admin.command("ping")["ok"] == 1:
        return

    try:
        client = MongoClient(
            os.environ.get("DATABASE_URL", "mongodb://127.0.0.1:27017")
        )
        db_name = os.environ.get("DATABASE_NAME", "scatter-protocol-web")
        client.server_info()  # Check if the connection is successful
        db = client[db_name]
        return db
    except Exception as e:
        print("Unable to connect to database.")
        raise e
