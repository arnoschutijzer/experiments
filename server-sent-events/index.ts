import { Request, Response } from "express";

import express from "express";
import bodyParser from "body-parser";

const app = express();

app.use(bodyParser.json());

type Client = {
  id: string;
  response: Response;
};
type Fact = {
  info: string;
  source: string;
};

let clients: Client[] = [];
let facts: Fact[] = [];

function eventsHandler(request: Request, response: Response) {
  const headers = {
    "Content-Type": "text/event-stream",
    Connection: "keep-alive",
    "Cache-Control": "no-cache",
  };
  response.writeHead(200, headers);

  const data = `data: ${createEvent("initial", facts)}\n\n`;

  response.write(data);

  const clientId = `${Date.now()}`;

  const newClient = {
    id: clientId,
    response,
  };

  clients.push(newClient);

  request.on("close", () => {
    console.log(`${clientId} Connection closed`);
    clients = clients.filter((client) => client.id !== clientId);
  });
}

function createEvent(type, data) {
  return JSON.stringify({ type, data });
}

app.get("/events", eventsHandler);

function sendEventsToAll(newFact) {
  clients.forEach((client) =>
    client.response.write(`data: ${createEvent("new", newFact)} \n\n`),
  );
}

async function addFact(request, response, next) {
  const newFact = request.body;
  facts.push(newFact);
  response.json(newFact);
  return sendEventsToAll(newFact);
}

app.post("/fact", addFact);
app.use("/static", express.static("index.html"));

app.listen(3123, () => {
  console.log(`listening`);
});
