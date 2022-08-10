
## Zabbix Database Cleaner
<p float="left">
  <img src="https://img.shields.io/github/downloads/limanmys/zabbix-cleaner-go/total?color=blue&style=for-the-badge" width="100" />
  <img src="https://img.shields.io/github/languages/code-size/limanmys/zabbix-cleaner-go?color=blue&style=for-the-badge" width="115" /> 
  <img src="https://img.shields.io/github/stars/limanmys/zabbix-cleaner-go?color=yellow&style=for-the-badge" width="65" /> 
  <img src="https://img.shields.io/github/go-mod/go-version/limanmys/zabbix-cleaner-go?color=blue&style=for-the-badge" width="75" />
</p>

----
This service is optionally prepared with the aim of cleaning meaningless data kept in the Zabbix database.

It currently supports postgresql and mysql databases. Tested in Zabbix version 6.0.1.

### Deployment
----
Depending on your request, the relevant package is downloaded from the [link here](https://github.com/limanmys/zabbix-cleaner-go/releases).

For Debian based systems:

```
sudo apt install ./zabbix-cleaner-x64.deb
```

For RHEL based systems:

```
sudo yum install ./zabbix-cleaner-x64.rpm
```

After installation, the environment file is filled with appropriate data.

```
sudo nano /opt/zabbix-cleaner/.env
```

The content should look like the following :

```
DB_DRIVER=postgresql/mysql # Database type
DB_HOST=10.1.1.1 # Database ip address
DB_PORT=5432/3306 # Database port
DB_NAME=zabbix # Zabbix database name (default:zabbix)
DB_USER=zabbix # Zabbix database username (default:zabbix)
DB_PASS= # Zabbix username's password
INTERVAL=30 # Data day to be protected
```

Finally, the service must be restarted.

```
sudo systemctl restart zabbix-cleaner
```

To make sure it works correctly :

```
sudo systemctl status zabbix-cleaner
```