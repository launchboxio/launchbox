import {Container, Row, Col, Card, Text} from '@nextui-org/react';

export default function ApplicationsList({data}) {
    console.log(data)
  return (
      <Container>
          {data.applications.map((item, index) => {
              return (
                  <Card color={"primary"}>
                      <Row justify={"center"} align={"center"}>
                          <Text h6 size={15} color="white" css={{ m: 0 }}>
                              {item.name}
                          </Text>
                      </Row>
                  </Card>
              )
          })}
      </Container>
  )
}

export async function getServerSideProps() {
    const res = await fetch('http://localhost:8080/applications')
    const data = await res.json()
    return { props: { data } }
}