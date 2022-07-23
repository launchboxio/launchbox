import NextAuth from "next-auth"
import GithubProvider from "next-auth/providers/github"
import OktaProvider from "next-auth/providers/okta"
import Auth0Provider from "next-auth/providers/auth0"
import GitlabProvider from "next-auth/providers/gitlab"
import OneLoginProvider from "next-auth/providers/onelogin"
import SequelizeAdapter from "@next-auth/sequelize-adapter"
import { Sequelize } from "sequelize"

var providers = []
if (process.env.GITHUB_ID) {
    providers.push(GithubProvider({
        clientId: process.env.GITHUB_ID,
        clientSecret: process.env.GITHUB_SECRET,
    }))
}

if (process.env.OKTA_CLIENT_ID) {
    providers.push(OktaProvider({
        clientId: process.env.OKTA_CLIENT_ID,
        clientSecret: process.env.OKTA_CLIENT_SECRET,
        issuer: process.env.OKTA_ISSUER
    }))
}

if (process.env.AUTH0_CLIENT_ID) {
    providers.push(Auth0Provider({
        clientId: process.env.AUTH0_CLIENT_ID,
        clientSecret: process.env.AUTH0_CLIENT_SECRET,
        issuer: process.env.AUTH0_ISSUER
    }))
}

if (process.env.GITLAB_CLIENT_ID) {
    providers.push(GitlabProvider({
        clientId: process.env.GITLAB_CLIENT_ID,
        clientSecret: process.env.GITLAB_CLIENT_SECRET
    }))
}

if (process.env.ONELOGIN_CLIENT_ID) {
    providers.push(OneLoginProvider({
        clientId: process.env.ONELOGIN_CLIENT_ID,
        clientSecret: process.env.ONELOGIN_CLIENT_SECRET,
    }))
}
const sequelize = new Sequelize(process.env.DATABASE_DSN)

export default NextAuth({
    providers: providers,
    session: {
        strategy: "jwt"
    },
    jwt: {
        encryption: true
    },
    pages: {
        // signIn: '/auth/signin',
        signOut: '/auth/signout',
        error: '/auth/error', // Error code passed in query string as ?error=
        verifyRequest: '/auth/verify-request', // (used for check email message)
        newUser: '/auth/new-user' // New users will be directed here on first sign in (leave the property out if not of interest)
    },
    adapter: SequelizeAdapter(sequelize),
})

