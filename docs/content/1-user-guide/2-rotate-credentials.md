---
title: "Rotating your credentials"
weight: 1
---

When you are using IAM Users you know that at some point you need to rotate your credentials. This is either ignored or forgotten.
In the cases where you do take the time to rotate your credentials, it's somewhat of a tedious task.

- You need to log in to the AWS Management Console.
- Locate your IAM User in the IAM service page.
- Generate new credentials.
- Download the CSV or copy and paste the values directly in your `~.aws/credentials` file.
- Validate that the new credentials are working.
- Mark the old credentials as in-active.
- Delete the old credentials.

As you can see these are a lot of actions, and potentially employees will not rotate their credentials.
You can execute all these steps in a single command:

```shell
aws-iam-user rotate --profile [Profile] --region [Region]
```
