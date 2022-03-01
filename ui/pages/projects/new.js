import {Button, Grid, Input} from '@nextui-org/react';
import {useState} from "react";
import {useRouter} from "next/router";

export default function Projectreate({data}) {
    const [name, setName] = useState("")
    const router = useRouter()
    const handleSubmit = async (event) => {
        event.preventDefault()
        const res = await fetch('http://localhost:8080/projects', {
            body: JSON.stringify({ name }),
            headers: {
                "Content-Type": "application/json"
            },
            method: "POST"
        })

        const data = await res.json()

        router.push({ pathname: `/projects/${data.id}` })
    }

    return (
        <>
            <h3>Creating new project</h3>
            <Grid.Container gap={2}>
                <form onSubmit={handleSubmit}>
                    <Input label="Project Name" placeholder="My awesome project" value={name} onChange={(evt) => {
                        setName(evt.target.value)
                    }}/>
                    <Button type={"submit"}>Submit</Button>
                </form>
            </Grid.Container>
        </>
    )
}
