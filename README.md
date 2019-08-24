# CodeBuildSlack

AWS CodeBuildの実行結果をLambdaを使ってSlackに通知する

## アーキテクチャ

![アーキテクチャ](./img/CodeBuildSlack.png)


## Cloudwatch Event (ルール)

- イベントパターン

```json
{
  "detail-type": [
    "CodeBuild Build State Change"
  ],
  "source": [
    "aws.codebuild"
  ],
  "detail": {
    "build-status": [
      "FAILED",
      "STOPPED",
      "SUCCEEDED",
      "IN_PROGRESS"
    ]
  }
}
```

## Lambda

serverless.ymlに記載

## Slack通知の結果

![Slack通知](./img/codebuild-slack-go2.png)