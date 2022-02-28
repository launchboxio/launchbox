import {Button, Grid, Input} from '@nextui-org/react';

export default function ApplicationView({data}) {
    // console.log(data)
    // const router = useRouter()
    // const { id } = router.query
    //
    // return <p>Application: {id}</p>
    console.log(data)
    return (
        <>
            <h3>Creating new project</h3>
            <Grid.Container gap={2}>
                <Input label="Project Name" placeholder="My awesome project" />
                <Button type={"submit"}>Submit</Button>
            </Grid.Container>
        </>
    )
}
