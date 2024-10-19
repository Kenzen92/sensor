import { StatusBar } from 'expo-status-bar';
import { useEffect, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';

export default function App() {
  const [sensorData, setSensorData] = useState(null);

  function getSensorData() {
    const BASE_URI = 'http://raspberrypi:5000/readings';
    console.log("Getting network data");
    fetch(BASE_URI)
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json(); // Parse the response as JSON
      })
      .then(data => {
        console.log('Sensor Data:', ...data); // Handle the sensor data
        // You can update your UI or state here with the sensor data
      })
      .catch(error => {
        console.error('Error fetching sensor data:', error);
      });      
  }

  useEffect(() => {
    getSensorData();
  }, [])
  return (
    <View style={styles.container}>
      <Text>Open up App.js to start working on your app!</Text>
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
