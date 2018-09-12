import click
import pickle
import client
import json
import os
    
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
        print('id: {} name: {} git_url: {} webhook_secret: {}'.format(b['id'], b['name'], b['git_url'], b['webhook_secret']))


@click.command()
@click.argument('name', nargs=1)
@click.argument('git_url', nargs=1)
@click.pass_context
def create_bot(ctx, name, git_url):
    c = ctx.obj['c']
    c.create_bot(name, git_url)

cli.add_command(list_bot)
cli.add_command(create_bot)
    
if not os.path.isfile('config.pk'):
    init()
else:
    with open('config.pk', 'rb') as f:
        conf = pickle.load(f)
        c = client.Client(conf['addr'], conf['password'])
        c.create_cookie()
        cli(obj={'c':c})

def init():
    print('server addr:')
    addr = input()
    print('input password:')
    password = input()
    print('new? (y)')
    new = input()
    if new == 'y':
        cli = client.Client(addr, password)
        cli.create_account()

    conf = {
        'addr': addr,
        'password': password
    }
    with open('config.pk', 'wb') as f:
        pickle.dump(conf, f)