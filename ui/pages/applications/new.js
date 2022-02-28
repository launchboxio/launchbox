import {Button, Grid, Input} from '@nextui-org/react';
import {useState} from "react";
import {useRouter} from "next/router";

export default function ApplicationCreate({data}) {
    const [name, setName] = useState("")
    const router = useRouter()
    const handleSubmit = async (event) => {
        event.preventDefault()
        const res = await fetch('http://localhost:8080/applications', {
            body: JSON.stringify({ name }),
            headers: {
                "Content-Type": "application/json"
            },
            method: "POST"
        })

        const data = await res.json()
        console.log(data)

        router.push({ pathname: `/applications/${data.id}` })
    }

    console.log(name)
    return (
        <>
            <h3>Creating new project</h3>
            <Grid.Container gap={2}>
                <form onSubmit={handleSubmit}>
                    <Input label="Application Name" placeholder="My awesome app" value={name} onChange={(evt) => {
                        setName(evt.target.value)
                    }}/>
                    <Button type={"submit"}>Submit</Button>
                </form>
            </Grid.Container>
        </>
    )
}
