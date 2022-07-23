import { getProviders } from "next-auth/react"
import {Button, Grid, Link} from "@nextui-org/react";

export default function Signin({providers}) {
    var buttons = Object.keys(providers).map((item, index) => {
        return (
            <Link href={providers[item].signinUrl} key={providers[item].name}>
                Login with { providers[item].name }
            </Link>
        )
    })
    return (
        <>
            <Grid.Container gap={2}>
                <h2>Sign In</h2>
                {buttons}
            </Grid.Container>
        </>
    )
}

export async function getServerSideProps() {
    var providers = await getProviders()
    return { props: { providers }}
}
