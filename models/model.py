import psycopg2
from dotenv import dotenv_values

def getName(conn):
    sql = """
    SELECT name FROM users"""

    with conn.cursor() as cur:
        cur.execute(sql)
        return [r[0] for r in cur.fetchall()]
    
def main():
    config = dotenv_values(".env")

    params = {
        "host": config["DB_HOST"],
        "database": config["DB_NAME"],
        "user": config["DB_USER"],
        "password": config["DB_PASSWORD"],
        "port": int(config["DB_PORT"])
    }

    with psycopg2.connect(**params) as conn:
        print(getName(conn))

if __name__ == '__main__':
    main()