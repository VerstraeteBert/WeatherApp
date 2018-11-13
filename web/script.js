const app = new Vue({
    el: "#app",
    data: {
        readings: [],
    },
    methods: {
        doReading() {
            this.readings.unshift({
                date: new Date(),
                temperature: Math.floor(Math.random() * 60 - 30),
                humidity: Math.floor(Math.random() * 50 + 50),
                lumen: Math.floor(Math.random() * 10 + 1500)
            })
        },
        getAllReadings() {
            let readings = [];
            for (let i = 0; i < 10; i++) {
                readings.push({
                    date: new Date(),
                    temperature: Math.floor(Math.random() * 60 - 30),
                    humidity: Math.floor(Math.random() * 50 + 50),
                    lumen: Math.floor(Math.random() * 10 + 1500)
                })
            }
            this.readings = readings;
        }
    },
});