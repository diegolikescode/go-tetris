import { Request, Response, Router } from "express";
import { uploadFileConfig } from "./config.js";
import { saveCsvToSqlite } from "./saveCsvToSqlite.js";
import { Op } from "sequelize";
import { fileURLToPath } from "url";
import { dirname } from "path";
import User from "./entities/User.js";

const appRouter = Router();
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

appRouter.post(
	"/files",
	uploadFileConfig.single("sweetCsv"),
	async (req: Request, res: Response) => {
		console.log("MY DIR DIR DIR", __dirname);
		const fileName = req.file.filename;
		const fullPath = `${__dirname}/../uploads/${fileName}`;
		await saveCsvToSqlite(fullPath);
		return res.status(201).json({ message: "all good" });
	}
);

appRouter.get("/users", async (req: Request, res: Response) => {
	const { query } = req;
	let usersResponse: { count: number; rows: any[] };
	if (Object.keys(query).length === 0) {
		usersResponse = await User.findAndCountAll();
	} else {
		const { q } = query;
		const searchableColumns = Object.keys(User.getAttributes()).filter(
			(colName) =>
				colName !== "id" &&
				colName !== "createdAt" &&
				colName !== "updatedAt"
		);

		const condition = searchableColumns.map((colName) => {
			const isLike = {};
			isLike[`${colName}`] = { [Op.like]: `%${q}%` };
			return isLike;
		});
		usersResponse = await User.findAndCountAll({
			where: { [Op.or]: [...condition] },
		});
	}
	return res.status(usersResponse.count > 0 ? 200 : 204).json(usersResponse);
});

export default appRouter;
