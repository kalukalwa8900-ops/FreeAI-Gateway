/*
Dependency install:
npm install openai

Run:
set OPENAI_API_KEY=your_key
node demo/demo_nodejs.js
*/

import OpenAI from "openai";

const client = new OpenAI({
  apiKey: process.env.OPENAI_API_KEY || "YOUR API KEY",
  baseURL: "https://api.chatanywhere.tech/v1",
});

// Non-stream response
async function gpt35Api(messages) {
  const completion = await client.chat.completions.create({
    model: "gpt-3.5-turbo",
    messages,
  });

  console.log(completion.choices[0]?.message?.content || "");
}

// Stream response
async function gpt35ApiStream(messages) {
  const stream = await client.chat.completions.create({
    model: "gpt-3.5-turbo",
    messages,
    stream: true,
  });

  for await (const chunk of stream) {
    const delta = chunk.choices?.[0]?.delta?.content;
    if (delta) {
      process.stdout.write(delta);
    }
  }
  process.stdout.write("\n");
}

(async () => {
  const messages = [
    { role: "user", content: "What is the relationship between Lu Xun and Zhou Shuren?" },
  ];

  // Non-stream call
  // await gpt35Api(messages);

  // Stream call
  await gpt35ApiStream(messages);
})();
