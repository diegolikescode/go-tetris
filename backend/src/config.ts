import { Sequelize } from "sequelize";
import path from "path";
import multer from "multer";

export const database = new Sequelize({
	dialect: "sqlite",
	storage: "./database.sqlite",
});

const storageFileConfig = multer.diskStorage({
	destination: (req, file, callBack) => {
		callBack(null, "./uploads/");
	},
	filename: (req, file, callBack) => {
		callBack(
			null,
			`${file.fieldname}-${Date.now()}${path.extname(file.originalname)}`
		);
	},
});

export const uploadFileConfig = multer({ storage: storageFileConfig });
