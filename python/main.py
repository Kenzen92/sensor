import time
import board
import busio
from adafruit_bme680 import Adafruit_BME680_I2C
import sqlite3


def main():
    print("Running main")

    # Create I2C bus
    i2c = busio.I2C(board.SCL, board.SDA)

    # Create sensor object
    bme680 = Adafruit_BME680_I2C(i2c)

    # Get the database and start filling in values
    connection = sqlite3.connect('sensor.db')
    cursor = connection.cursor()

    cursor.execute('''
    CREATE TABLE IF NOT EXISTS environmental_readings (
        id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique ID for each record
        temperature REAL NOT NULL,             -- Temperature in Celsius
        humidity REAL NOT NULL,                -- Humidity in percentage
        pressure REAL NOT NULL,                -- Pressure in hPa
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP  -- Timestamp of the reading
    );
    ''')

    # Main loop to print sensor readings
    while True:
        cursor.execute(f'''
        INSERT INTO environmental_readings (temperature, humidity, pressure)
        VALUES ({bme680.temperature:.2f}, {bme680.humidity:.2f}, {bme680.pressure:.2f});
        ''')
        connection.commit()
        time.sleep(5)


if __name__ == '__main__':
    main()