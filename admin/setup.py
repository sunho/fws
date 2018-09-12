from setuptools import setup

setup(
    name='fwsadm',
    version=1.0,
    packages=['fwsadm'],
    install_requires=['click', 'requests'],
    entry_points={'console_scripts': ['fwsadm = fwsadm.fwsadm:main']}
)