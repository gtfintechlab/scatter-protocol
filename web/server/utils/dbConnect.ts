import mongoose from "mongoose";

async function dbConnect(): Promise<void> {
    if (mongoose.connections[0].readyState) return;

    await mongoose
        .connect(process.env.DATABASE_URL ?? "mongodb://127.0.0.1:27017" as string, {
            socketTimeoutMS: 360000,
            dbName: process.env.DATABASE_NAME ?? "scatter-protocol",
        })
        .catch((error) => {
            console.error("Unable to connect to database.");

            throw error;
        });
}
export default dbConnect;