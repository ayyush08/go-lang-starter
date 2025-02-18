import express from 'express';

const app = express();
const PORT = 5000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/get', (req, res) => {
    res.send('This is a GET request response');
});


app.post('/post', (req, res) => {
    res.status(201).send(req.body)
})

app.post('/postform', (req, res) => {
    console.log(req.body);
    
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});

export default app;