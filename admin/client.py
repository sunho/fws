import requests
import json
session = requests.session()

class Client:
    def __init__(self, addr, password):
        self.addr = addr
        self.sess = requests.session()
        self.password = password

    def create_account(self):
        data = {
            'username': 'admin',
            'nickname': 'admin',
            'key': 'admin',
            'password': self.password
        }
        res = self.sess.post(self.addr + '/api/register', data=json.dumps(data))
        if res.status_code != 201:
            raise Exception('creating account failed')

    def create_cookie(self):
        data = {
            'username': 'admin',
            'password': self.password
        }
        res = self.sess.post(self.addr + '/api/login', data=json.dumps(data))
        if res.status_code != 201:
            raise Exception('creating cookie failed')

    def create_bot(self, name, git_url):
        data = {
            'name': name,
            'git_url': git_url
        }
        res = self.sess.post(self.addr + '/api/admin/bot', data=json.dumps(data))
        if res.status_code != 201:
            raise Exception('creating bot failed')
    
    def list_bot(self):
        res = self.sess.get(self.addr + '/api/admin/bot')
        if res.status_code != 200:
            raise Exception('listing bot failed')
        
        return res.json()
