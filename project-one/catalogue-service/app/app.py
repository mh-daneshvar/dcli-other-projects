from flask import Flask, request, jsonify
import mysql.connector
import os

app = Flask(__name__)

# MySQL connection settings
def get_db_connection():
    connection = mysql.connector.connect(
        host="mysql_db",
        user=os.getenv("MYSQL_USER"),
        password=os.getenv("MYSQL_PASSWORD"),
        database=os.getenv("MYSQL_DATABASE")
    )
    return connection

# API endpoint to add an item
@app.route('/items', methods=['POST'])
def add_item():
    data = request.json
    name = data['name']

    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("INSERT INTO items (name) VALUES (%s)", (name,))
    conn.commit()
    cursor.close()
    conn.close()

    return jsonify({"message": "Item added!"}), 201

# API endpoint to get all items
@app.route('/items', methods=['GET'])
def get_item():
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM items")
    items = cursor.fetchall()
    cursor.close()
    conn.close()

    return jsonify(items)

# Main entry point
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
