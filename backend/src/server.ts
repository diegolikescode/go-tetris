import "reflect-metadata";
import express from "express";
import CORS from "cors";
import router from "./routes.js";
import { database } from "./config.js";
const app = express();
const port = 3069;

(async () => {
	try {
		await database.sync();
	} catch (err) {
		console.log("error in the database sync", err);
	}
})();

app.use(express.json());
app.use(
	CORS({
		credentials: false,
		origin: "*",
		methods: ["GET", "POST", "PUT", "DELETE", "PATCH"],
		allowedHeaders: [
			"Access-Control-Allow-Origin",
			"X-CSRF-Token",
			"X-Requested-With",
			"Accept",
			"Accept-Version",
			"Content-Length",
			"Content-MD5",
			"Content-Type",
			"Date",
			"X-Api-Version",
			"Authorization",
		],
		optionsSuccessStatus: 200, // legacy stuff (cors' docs)
	})
);

app.use("/api", router);
app.listen(port, () => console.log(`running at ${port}`));
