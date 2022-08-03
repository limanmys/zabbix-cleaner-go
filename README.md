## Zabbix Database Cleaner
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
INTERVAL=60 # Data day to be protected
SAVE_DELETED_ROWS=true/false # If true, service will save your rows under /var/log/zabbix-cleaner
```

Finally, the service must be restarted.

```
sudo systemctl restart zabbix-cleaner
```

To make sure it works correctly :

```
sudo systemctl status zabbix-cleaner
```