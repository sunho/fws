import requests
import json
session = requests.session()

def res_error(res):
    return Exception('{}:{}'.format(res.status_code, res.text))

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
        if not res.ok:
            raise res_error(res)

    def create_cookie(self):
        data = {
            'username': 'admin',
            'password': self.password
        }
        res = self.sess.post(self.addr + '/api/login', data=json.dumps(data))
        if not res.ok:
            raise res_error(res)

    def create_userbot(self, username, bot_id):
        data = {
            'id': bot_id
        }
        res = self.sess.post(self.addr + '/api/admin/user/{}/bot'.format(username), data=json.dumps(data))
        if not res.ok:
            raise res_error(res)

    def list_userbot(self, username):
        res = self.sess.get(self.addr + '/api/admin/user/{}/bot'.format(username))
        if not res.ok:
            raise res_error(res)

        return res.json()
        

    def delete_userbot(self, username, bot_id):
        res = self.sess.delete(self.addr + '/api/admin/user/{}/bot/{}'.format(username, bot_id))
        if not res.ok:
            raise res_error(res)

    def list_bot(self):
        res = self.sess.get(self.addr + '/api/admin/bot')
        if not res.ok:
            raise res_error(res)
        
        return res.json()

    def create_bot(self, name, git_url):
        data = {
            'name': name,
            'git_url': git_url
        }
        res = self.sess.post(self.addr + '/api/admin/bot', data=json.dumps(data))
        if not res.ok:
            raise res_error(res)

    def put_bot(self, bot):
        res = self.sess.put(self.addr + '/api/admin/bot/{}'.format(bot['id']), data=json.dumps(bot))
        if not res.ok:
            raise res_error(res)

    def delete_bot(self, id):
        res = self.sess.delete(self.addr + '/api/admin/bot/{}'.format(id))
        if not res.ok:
            raise res_error(res)
            
    def list_invite(self):
        res = self.sess.get(self.addr + '/api/admin/invite')
        if not res.ok:
            raise res_error(res)
        
        return res.json()

    def create_invite(self, username):
        data = {
            'username': username,
        }
        res = self.sess.post(self.addr + '/api/admin/invite', data=json.dumps(data))
        if not res.ok:
            raise res_error(res)
        
        return res.text
    
    def delete_invite(self, username):
        res = self.sess.delete(self.addr + '/api/admin/invite/{}'.format(username))
        if not res.ok:
            raise res_error(res)
            
    def list_user(self):
        res = self.sess.get(self.addr + '/api/admin/user')
        if not res.ok:
            raise res_error(res)

        return res.json()
        
    def delete_user(self, username):
        res = self.sess.delete(self.addr + '/api/admin/user/{}'.format(username))
        if not res.ok:
            raise res_error(res)