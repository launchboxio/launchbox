import {Grid, Button, Row, Col, Card, Text, Divider, Link} from '@nextui-org/react';

export default function ApplicationsList({data}) {
  return (
      <Grid.Container gap={2}>
          {data.applications.map((item, index) => {
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
                                  <Link href={`applications/${item.id}`}>View</Link>
                              </Row>
                          </Card.Footer>
                      </Card>
                  </Grid>
              )
          })}
          <Link href={"/applications/new"}>
              New Application
          </Link>
      </Grid.Container>
  )
}

export async function getServerSideProps() {
    const res = await fetch('http://localhost:8080/applications')
    const data = await res.json()
    return { props: { data } }
}