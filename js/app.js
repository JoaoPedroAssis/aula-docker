const express = require('express');
const { Pool } = require('pg');

const app = express();
const port = 3000;

// Configuração de conexão com o banco de dados PostgreSQL
const pool = new Pool({
  user: 'postgres',
  host: 'localhost',
  database: 'postgres',
  password: 'postgres',
  port: 5432,
});

// Rota "Hello World"
app.get('/', (req, res) => {
  res.send('Hello, World!');
});

// Rota para testar a conexão com o banco de dados
app.get('/test-db-connection', async (req, res) => {
  try {
    const client = await pool.connect();
    client.release(); // Libera o cliente para o pool após a conexão ser estabelecida com sucesso
    res.status(200).json({ message: 'Conexão com o banco de dados estabelecida com sucesso!' });
  } catch (error) {
    console.error('Erro ao conectar ao banco de dados:', error);
    res.status(500).json({ message: 'Erro ao conectar ao banco de dados.' });
  }
});

// Inicie o servidor na porta especificada
app.listen(port, () => {
  console.log(`Servidor Express rodando na porta ${port}`);
});
