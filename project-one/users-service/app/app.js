const express = require('express');
const mongoose = require('mongoose');
const app = express();

// Middleware
app.use(express.json());

// Connect to MongoDB
mongoose.connect('mongodb://mongo_db:27017/mydb', {
    useNewUrlParser: true,
    useUnifiedTopology: true
})
    .then(() => console.log('MongoDB connected'))
    .catch(err => console.log(err));

// Define a simple schema and model
const ItemSchema = new mongoose.Schema({ name: String });
const Item = mongoose.model('Item', ItemSchema);

// API endpoint to add an item to the database
app.post('/items', async (req, res) => {
    const newItem = new Item({ name: req.body.name });
    await newItem.save();
    res.send(newItem);
});

// API endpoint to get all items
app.get('/items', async (req, res) => {
    const items = await Item.find();
    res.send(items);
});

// Start the server
const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Users Service is running on port ${PORT}`);
});
