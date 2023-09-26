from flask import Flask, jsonify
import psycopg2

app = Flask(__name__)

# Configuração de conexão com o banco de dados PostgreSQL
# db_config = {
#     'user': 'seu_usuario',
#     'password': 'sua_senha',
#     'host': 'localhost',
#     'port': 5432,
#     'database': 'seu_banco_de_dados'
# }

# Rota "Hello World"
@app.route('/')
def hello_world():
    return 'Hello, World!'

# Rota para testar a conexão com o banco de dados
# @app.route('/test-db-connection')
# def test_db_connection():
#     try:
#         conn = psycopg2.connect(**db_config)
#         conn.close()
#         return jsonify({'message': 'Conexão com o banco de dados estabelecida com sucesso!'})
#     except Exception as e:
#         print('Erro ao conectar ao banco de dados:', e)
#         return jsonify({'message': 'Erro ao conectar ao banco de dados.'}), 500

if __name__ == '__main__':
    app.run(debug=True)
