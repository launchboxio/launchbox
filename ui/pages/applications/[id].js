import {Grid, Button, Row, Col, Card, Text, Divider, Link} from '@nextui-org/react';

export default function ApplicationView({data}) {
    // console.log(data)
    // const router = useRouter()
    // const { id } = router.query
    //
    // return <p>Application: {id}</p>
    console.log(data)
    let content = <h4>You don't have any projects for this app. <a href={`/projects/new?application_id=${data.application.id}`}>Create one now</a></h4>
    if (data.projects.projects.length > 0) {
        content = data.projects.projects.map((item, index) => {
            return (
                <Grid xs={12}>
                    <Card key={item.id}>
                        <Card.Header>
                            {item.name}
                        </Card.Header>
                        <Divider />
                        <Card.Body css={{ py: '$10'}}>
                            <Text>Lorem ipsum</Text>
                        </Card.Body>
                        <Divider />
                        <Card.Footer>
                            <Row justify={"flex-end"}>
                                <Button.Group>
                                    <Button>Logs</Button>
                                    <Button>Shell</Button>
                                    <Button>Actions</Button>
                                </Button.Group>
                            </Row>
                        </Card.Footer>
                    </Card>
                </Grid>
            )
        })
    }
    return (
        <Grid.Container gap={2}>
            <h2>{data.application.name}</h2>
            <Row>
                <h3>Projects</h3>
            </Row>
            {content}
        </Grid.Container>
    )
}

export async function getServerSideProps(context) {
    const data = {}
    const app = await fetch(`http://localhost:8080/applications/${context.params.id}`)
    const projects = await fetch(`http://localhost:8080/projects?application_id=${context.params.id}`)
    data.application = await app.json()
    data.projects = await projects.json()
    return { props: { data } }
}