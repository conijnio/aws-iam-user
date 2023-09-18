---
title: "Retrieve credential context"
weight: 1
---

When you have multiple sets of credentials in your `~.aws/credentials` file. It can become harder to know what credentials belong to what AWS Account.
You can use a descriptive profile name, but what if you did not do that? How do you know to what AWS Account the credentials belong?

With the following command you will find out to what AWS Account your credentials belong too:

```shell
aws-iam-user --profile [Profile] --region [Region]
```
