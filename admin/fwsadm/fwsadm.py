from .client import Client

import click
import pickle
import json
import os

def print_item(i):
    t = ''
    for k, v in i.items():
        t += '{}:{} '.format(k, v)
    print(t)

def edit_item(i):
    for k, v in i.items():
        print('{}:{}'.format(k,v))
        n = input('edit? (y)')
        if n == 'y':
            if type(i[k]) is str:
                i[k] = input('new value:')
            else:
                i[k] = eval(input('new value:'))
    print(i)

@click.command()
@click.pass_context
def list_user(ctx):
    c = ctx.obj['c']
    users = c.list_user()
    for u in users:
        print_item(u)
        bots = c.list_userbot(u['username'])
        t = ['{}:{}'.format(x['id'], x['name']) for x in bots]
        print('owned bots: ' + ' '.join(t))
    
@click.command()
@click.argument('username', nargs=1)
@click.pass_context
def delete_user(ctx, username):
    c = ctx.obj['c']
    c.delete_user(username)

@click.command()
@click.argument('username', nargs=1)
@click.argument('bot_id', type=int, nargs=1)
@click.pass_context
def give_bot(ctx, username, bot_id):
    c = ctx.obj['c']
    c.create_userbot(username, bot_id)

@click.command()
@click.argument('username', nargs=1)
@click.argument('bot_id', type=int, nargs=1)
@click.pass_context
def ungive_bot(ctx, username, bot_id):
    c = ctx.obj['c']
    c.delete_userbot(username, bot_id)
    
@click.group()
@click.pass_context
def cli(ctx):
    pass

@click.command()
@click.pass_context
def list_bot(ctx):
    c = ctx.obj['c']
    bots = c.list_bot()
    for b in bots:
        print_item(b)

@click.command()
@click.argument('name', nargs=1)
@click.argument('git_url', nargs=1)
@click.pass_context
def create_bot(ctx, name, git_url):
    c = ctx.obj['c']
    c.create_bot(name, git_url)

@click.command()
@click.argument('nameorid', nargs=1)
@click.argument('value', nargs=1)
@click.pass_context
def edit_bot(ctx, nameorid, value):
    c = ctx.obj['c']
    bots = c.list_bot()
    b = []
    if nameorid == 'id':
        b = [x for x in bots if x['id'] == value]
    elif nameorid == 'name':
        b = [x for x in bots if x['name'] == value]
    else:
        print('invalid input')
        return
    
    edit_item(b[0])
    c.put_bot(b[0])

@click.command()
@click.argument('nameorid', nargs=1)
@click.argument('value', nargs=1)
@click.pass_context
def delete_bot(ctx, nameorid, value):
    c = ctx.obj['c']
    if nameorid == 'id':
        c.delete_bot(value)
    elif nameorid == 'name':
        bots = c.list_bot()
        matches = [x for x in bots if x['name'] == value]
        if len(matches) == 0:
            print('no such bot')
            return
        for m in matches:
            c.delete_bot(m['id'])
    else:
        print('invalid input')

@click.command()
@click.pass_context
def list_invite(ctx):
    c = ctx.obj['c']
    invites = c.list_invite()
    for i in invites:
        print_item(i)
        print(c.addr + '/invite?key=' + i['key'] + '&username=' + i['username'])

@click.command()
@click.argument('username', nargs=1)
@click.pass_context
def create_invite(ctx, username):
    c = ctx.obj['c']
    key = c.create_invite(username)
    print('key: ' + key)
    print(c.addr + '/invite?key=' + key + '&username=' + username)

@click.command()
@click.argument('username', nargs=1)
@click.pass_context
def delete_invite(ctx, username):
    c = ctx.obj['c']
    key = c.delete_invite(username)
    
cli.add_command(list_user)
cli.add_command(delete_user)

cli.add_command(list_bot)
cli.add_command(create_bot)
cli.add_command(edit_bot)
cli.add_command(delete_bot)

cli.add_command(list_invite)
cli.add_command(create_invite)
cli.add_command(delete_invite)

cli.add_command(give_bot)
cli.add_command(ungive_bot)
    
def main():
    if not os.path.isfile('config.pk'):
        init()
    else:
        with open('config.pk', 'rb') as f:
            conf = pickle.load(f)
            c = Client(conf['addr'], conf['password'])
            c.create_cookie()
            cli(obj={'c':c})

def init():
    addr = input('server addr:')
    password = input('input password:')
    new = input('new? (y)')
    if new == 'y':
        cli = Client(addr, password)
        cli.create_account()

    conf = {
        'addr': addr,
        'password': password
    }
    with open('config.pk', 'wb') as f:
        pickle.dump(conf, f)

if __name__ == '__main__':
    main()