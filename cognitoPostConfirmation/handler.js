var aws = require('aws-sdk');

exports.handler = async (event, context, callback) => {
  const documentClient = new aws.DynamoDB.DocumentClient({});
  console.log("function start")
  if (event.request.userAttributes.sub) {
    console.log("event.request.userAttributes.sub exists")
    const params = {
      User: event.request.userAttributes.sub,
      StartWork: "user"
    }
    try {
      await documentClient.put({ TableName: process.env.TABLE_NAME, Item: params }).promise()
      console.log("put success")
    } catch (error) {
      console.log(error)
      console.log("error!!")
    }
    callback(null, event);
  } else {
    console.log("not exists")
    callback(null, event);
  }
};
