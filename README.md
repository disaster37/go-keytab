# go-keytab
Handle the keytab file from command line

It's a king of wrapper on kutil, klist command. So you need to have this 2 binary.
You can get them with krb5-workstation (redhat) or krb5-user (debian) package.

You can use it to:
 - add principal on keytab
 - delete kaytab file
 - check if principal exist on keytab

## Usage

### Global parameters

You need to set the full path of keytab file for each command.

```sh
./go-keytab --path /etc/security/keytab/service.keytab ... 
```

You need to specify the following parameters:
- **path**: The full path for the keytab file.


### Add one principal with one cipher on keytab

You need to lauch the following command:

```sh
./go-keytab --path /etc/security/keytab/service.keytab add-principal --principal host/service@DOMAIN.COM --password my_long_password --cipher aes256-cts-hmac-sha1-96
```

You need to specify the following parameters:
- **--principal**: The principal to add on keytab. You must suffix it with the realm.
- **--password** : The password associated to principal.
- **--cipher** (optionnal): The cipher to encrypt password. Default to `aes256-cts-hmac-sha1-96`.

It return error if principal/cipher already exist on keytab.

### Add one principal with multiple ciphers on keytab

You need to lauch the following command:

```sh
./go-keytab --path /etc/security/keytab/service.keytab add-principal --principal host/service@DOMAIN.COM --password my_long_password --ciphers arcfour-hmac,des-cbc-md5,des3-cbc-sha1,aes256-cts-hmac-sha1-96,aes128-cts-hmac-sha1-96
```

You need to specify the following parameters:
- **--principal**: The principal to add on keytab. You must suffix it with the realm.
- **--password** : The password associated to principal.
- **--ciphers** (optionnal): The cipher list to encrypt password. Default to `aes256-cts-hmac-sha1-96`.


It return error if one of them principal/cipher already exist on keytab

If you need to generate keytab for AD, you can use the following cyphers:

```
--ciphers arcfour-hmac,des-cbc-md5,aes128-cts-hmac-sha1-96,des3-cbc-sha1,aes256-cts-hmac-sha1-96
```

### Check if principal exist on keytab

You need to lauch the following command:

```sh
./go-keytab --path /etc/security/keytab/service.keytab check-principal --principal host/service@DOMAIN.COM --cipher aes256-cts-hmac-sha1-96
```

You need to specify the following parameters:
- **--principal**: The principal to check on keytab. You must suffix it with the realm.
- **--cipher** (optionnal): The cipher to encrypt password. Default to `aes256-cts-hmac-sha1-96`.



It return exit code `1` if principal is not found on keytab with the provided cipher.

### Delete keytab file

You need to lauch the following command:

```sh
./go-keytab --path /etc/security/keytab/service.keytab delete-keytab
```