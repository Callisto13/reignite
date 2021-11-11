#!/usr/bin/env python3

from test import Test
from metal import Welder
from config import Config
import click
import os
import sys
import random
import string
from os.path import dirname, abspath

@click.group()
def cli():
    """
    General thing doer for flintlock
    """
    pass

@cli.command()
@click.option('-c', '--config-file', type=str, help='Path to configuration file')
@click.option('-o', '--org-id', type=str, help='Equinix organisation id (required)')
@click.option('-p', '--project-name', type=str, help='Name of the project to create (default: randomly generated)')
@click.option('-k', '--ssh-key-name', type=str, help='Name of the ssh key to create and attach to the device (default: randomly generated)')
@click.option('-d', '--device-name', type=str, help='Name of the device to create (default: randomly generated)')
@click.option('-s', '--skip-teardown', is_flag=True, help='Skip cleanup of Equinix infrastructure for debugging after run')
def run_e2e(config_file, org_id, project_name, ssh_key_name, device_name, skip_teardown):
    token = os.environ.get("METAL_AUTH_TOKEN")
    if token is None:
        click.echo("must set METAL_AUTH_TOKEN")
        sys.exit()

    cfg = Config()
    if config_file is not None:
        try:
            cfg.load_config_from_file(config_file)
        except ValueError as e:
            click.echo(str(e))
            sys.exit()

    cfg.set_run_flag_config(org_id, project_name, ssh_key_name, device_name, skip_teardown)
    try:
        cfg.validate_run()
    except ValueError as e:
        click.echo(str(e))
        sys.exit()

    if cfg['device']['id'] == None:
        click.echo(f"Running e2e tests. Will create project `{cfg['project_name']}`, ssh_key `{cfg['device']['ssh_key_name']}` and device `{cfg['device']['name']}`")
        click.echo("Note: this will create and bootstrap a new device in Equinix and may take some time")
    else:
        tool_dir = dirname(abspath(__file__))
        if os.path.exists(tool_dir+"/private.key") != True:
            click.echo(f"`private.key` file must be saved at `{tool_dir}` when `--existing-device-id` flag set")
            sys.exit()
        click.echo("running e2e tests using device `{cfg['device']['id']}`")

    runner = Test(token, cfg)
    with runner:
        runner.setup()
        runner.run_tests()

    if cfg['test']['skip_teardown']:
        dev_id, dev_ip = runner.device_details()
        click.echo(f"Device `{dev_id}` left alive for debugging. Use with `--config-file` setting `device.id` to re-run tests. SSH command `ssh -i test/tools/private.key root@{dev_ip}`. Teardown device with `teardown-device` command.")

@cli.command()
@click.option('-o', '--org-id', type=str, help='Equinix organisation id (required)')
@click.option('-p', '--project-id', type=str, help='ID of the project to create the device in (required)')
@click.option('-k', '--ssh-key-name', type=str, help='Name of the ssh key to create and attach to the device (default: randomly generated)', default=generated_key_name())
@click.option('-d', '--device-name', type=str, help='Name of the device to create (default: randomly generated)', default='T-800')
@click.option('-u', '--userdata', type=str, help='String containing shell bootstrap userdata (default: standard flintlockd bootstrapping, see readme for details)')
def create_device(org_id, project_id, ssh_key_name, device_name, userdata):
    token = os.environ.get("METAL_AUTH_TOKEN")
    if token is None:
        click.echo("must set METAL_AUTH_TOKEN")
        sys.exit()
    if org_id is None:
        click.echo("must set --org-id")
        sys.exit()
    if project_id is None:
        click.echo("must set --project-id")
        sys.exit()

    click.echo(f"Creating device {device_name}")

    welder = Welder(token, org_id)
    ip = welder.create_all(project_id, device_name, ssh_key_name, userdata)

    click.echo(f"Device {device_name} created. SSH command `ssh -i hack/tools/private.key root@{ip}`. Run tests with `run-e2e`. Delete with `delete-device`.")

@cli.command()
@click.option('-o', '--org-id', type=str, help='Equinix organisation id (required)')
@click.option('-d', '--device-id', type=str, help='Name of the device to delete (required)')
def delete_device(org_id, device_id):
    token = os.environ.get("METAL_AUTH_TOKEN")
    if token is None:
        click.echo("must set METAL_AUTH_TOKEN")
        sys.exit()
    if org_id is None:
        click.echo("must set --org-id")
        sys.exit()
    if device_id is None:
        click.echo("must set --device-id")
        sys.exit()

    click.echo(f"Deleting device {device_id}")

    welder = Welder(token, org_id)
    welder.delete_device(device_id)

if __name__ == "__main__":
    cli()
