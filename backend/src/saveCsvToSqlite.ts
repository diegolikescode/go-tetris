import { createReadStream, unlinkSync } from "fs";
import csv from "fast-csv";
import User from "./entities/User.js";

export const saveCsvToSqlite = async (csvUrl: string) => {
	const stream = createReadStream(csvUrl);
	const collectionCsv: string[][] = [];
	const csvFileStream = csv
		.parse()
		.on("data", (data) => {
			collectionCsv.push(data);
		})
		.on("end", async () => {
			const csvHeader = collectionCsv[0];
			collectionCsv.splice(0, 1);

			const formattedForModel = collectionCsv.map((row) => {
				const userData = {};
				csvHeader.forEach((tuple, idx) => {
					userData[`${tuple}`] = row[idx];
				});
				return userData;
			});

			await User.bulkCreate(formattedForModel);

			// the line bellow will destroy the csv at the end.
			// unlinkSync(csvUrl);
		});
	stream.pipe(csvFileStream);
};
