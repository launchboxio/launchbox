import { getToken } from "next-auth/jwt"

const secret = process.env.NEXTAUTH_SECRET

export default async (req, res) => {
    const token = await getToken({
        req,
        secret,
        encryption: true,
        raw: true
    });

    console.log(req.constructor.name)
    req.headers.authorization = `Bearer ${token}`;
    req.url = req.url.replace(/^\/api/, "");
    req.host = "http://localhost:8080"
    console.log(req.url)
    console.log(req.hostname)

}