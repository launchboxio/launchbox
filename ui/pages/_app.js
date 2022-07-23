import '../styles/globals.css'
import { NextUIProvider } from '@nextui-org/react';
import Layout from "../layouts/base";
import { SessionProvider } from "next-auth/react"

export default function App({
    Component,
    pageProps: { session, ...pageProps }
}) {
    return (
        <NextUIProvider>
            <Layout>
                <SessionProvider session={session}>
                    <Component {...pageProps} />
                </SessionProvider>
            </Layout>
        </NextUIProvider>
    )
}

