/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

console.log(`-----MY MONGO-INIT.JS SCRIPT EXECUTED-----`)


db.createUser({
    user: "southperry",
    pwd: "southperry",
    roles: [
        {
            role: "readWrite",
            db: "southperry",
        },
    ],

});