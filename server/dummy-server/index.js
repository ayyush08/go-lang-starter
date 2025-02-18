import express from 'express';

const app = express();
const PORT = 5000;

app.use(express.json());

app.get('/get', (req, res) => {
    res.send('This is a GET request response');
});

app.post('/postform', (req, res) => {
    res.send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});

export default app;