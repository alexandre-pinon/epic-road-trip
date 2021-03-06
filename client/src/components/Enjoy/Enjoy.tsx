import { AspectRatio, Avatar, Button, Card, Container, createStyles, Title, Group, Space, Tooltip, Text, Image, Paper, Grid, Center, SimpleGrid, ActionIcon } from "@mantine/core";
import axios from "axios";
import { useState } from "react";
import { ArrowBackUp, ArrowForwardUp, Bed, Bike, Car, PlaneInflight, Search, Train, Walk } from 'tabler-icons-react';
import { useNavigate } from "react-router-dom";

const useStyles = createStyles((theme) => ({
  button: {
    borderRadius: 0,

    '&:not(:first-of-type)': {
      borderLeftWidth: 0,
    },

    '&:first-of-type': {
      borderTopLeftRadius: theme.radius.sm,
      borderBottomLeftRadius: theme.radius.sm,
    },

    '&:last-of-type': {
      borderTopRightRadius: theme.radius.sm,
      borderBottomRightRadius: theme.radius.sm,
    },
  },

  cardEnjoy: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
  },

  sectionEnjoy: {
    borderBottom: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
      }`,
    paddingLeft: theme.spacing.md,
    paddingRight: theme.spacing.md,
    paddingBottom: theme.spacing.md,
  },

  likeEnjoy: {
    color: theme.colors.red[6],
  },

  labelEnjoy: {
    textTransform: 'uppercase',
    fontSize: theme.fontSizes.xs,
    fontWeight: 700,
  },
}));

export function Enjoy({ fulTrip }: any) {
  const { classes } = useStyles();
  const navigate = useNavigate();

  const goTravel = async () => {
    console.log("Go to travel page!")
    navigate('/travel');
  };

  const goSleep = async () => {
    console.log("Go to sleep page!")
    navigate('/sleep');
  };

  const [id, setId] = useState(0)
  const [city, setCity] = useState('')
  const [radius, setRadius] = useState('')
  const [name, setName] = useState('')
  const [rating, setRating] = useState('')
  const [vicinity, setVicinity] = useState('')
  const [icon, setIcon] = useState('')

  const [enjoy, setEnjoy] = useState([{
    id,
    name,
    rating,
    vicinity,
    icon
  }])

  console.log(fulTrip)

  let [selectedEnjoy, setSelectedEnjoy] = useState('')


  const [toggleEnjoy, setToggleEnjoy] = useState(false)

  const retrieveEnjoy = (event: any) => {
    axios.defaults.withCredentials = true
    event.preventDefault()
    let params = {
      city: fulTrip.startCity,
      constraints: {
        radius: 10000,
      }
    };
    console.log("event: ", params);

    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/enjoy',
      data: params,
    })
      .then((response) => {
        console.log(response.data);
        let id = 0
        let enjoyActivities: any = [];
        response.data.data.forEach((data: any) => {
          let activities = {
            id: id,
            name: data.name,
            rating: data.rating,
            vicinity: data.vicinity,
            icon: data.icon,
          }
          enjoyActivities.push(activities)
          id++;
        })
        setEnjoy(enjoyActivities)
      })
      .catch(function (error) {
        console.log(error);
      });

    setToggleEnjoy(!toggleEnjoy)

  };

  const showEnjoy = () => {
    setToggleEnjoy(!toggleEnjoy)
  }

  const saveEnjoy = () => {
    console.log(fulTrip)
  }

  const selectEnjoy = (id: number, type: string) => {

    setSelectedEnjoy(type)
    console.log(id)
    setId(id)

    if (type == "Enjoy") {
      fulTrip.setEnjoyName(enjoy[id])
    }

    console.log("fulTrip content: ", fulTrip)
    console.log("scroll to bottom: ", document.body.offsetHeight)
    window.scroll({
      top: document.body.offsetHeight,
      left: 0,
      behavior: 'smooth',
    });
  }

  return (
    <Container>
      <Card withBorder radius="md" p="md" className={classes.cardEnjoy}>

        {/* <form onSubmit={retrieveEnjoy}>
          <h1 className="h3 mb-3 fw-normal">Enjoy Activities</h1>

          <input type="text" className="form-control" placeholder="City" required
            onChange={e => setCity(e.target.value)}
          />

          <input type="text" className="form-control" placeholder="Radius" required
            onChange={e => setRadius(e.target.value)}
          />

          <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form> */}
        <Title
          align="center"
          data-testid="title"
          sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
        >
          Enjoy Activities


        </Title>
        <Space h="xl" />

        <Center>
          <Button size="xl" onClick={retrieveEnjoy} compact variant="subtle">
            ????
          </Button>
        </Center>
        <Space h="xs" />


        <SimpleGrid cols={2} spacing="md" breakpoints={[{ maxWidth: 'sm', cols: 1 }]}>
          {
            toggleEnjoy ? (
              enjoy.map((item) => (
                <Paper shadow="xl" p="md" withBorder key={item.id}>
                  <Center>
                    <Grid>
                      <Image
                        width={30}
                        height={60}
                        fit="contain"
                        radius="sm"
                        src={item.icon}
                      />
                    </Grid>
                  </Center>
                  <Space h="xs" />
                  <Grid><Text weight={700}>Name :  </Text> <Text> &nbsp; {item.name}</Text></Grid>
                  <Grid><Text weight={700}>Rating :  </Text> <Text> &nbsp; {item.rating} / 5</Text></Grid>
                  <Grid><Text weight={700}>Vicinity :  </Text> <Text> &nbsp; {item.vicinity}</Text></Grid>
                  <Space h="xl" />
                  <Center>
                    <ActionIcon onClick={() => selectEnjoy(item.id, 'Enjoy')} variant="outline">????</ActionIcon>
                  </Center>
                </Paper>
              ))
            )
              : null
          }
        </SimpleGrid>

        <div>
          {
            selectedEnjoy === "Enjoy" ? (
              <>
                <Space h="xl" />

                <Title
                  order={2}
                  align="center"
                  sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
                >
                  Enjoy Activities


                </Title>
                <Space h="xl" />                <Paper shadow="xl" p="md" withBorder >
                  {/* <Grid><Text weight={700}>Enjoy ID  {enjoy[id].id}</Text> </Grid> */}
                  <Center>
                    <Grid>
                      <Image
                        width={30}
                        height={60}
                        fit="contain"
                        radius="sm"
                        src={enjoy[id].icon}
                      />
                    </Grid>
                  </Center>
                  <Space h="xs" />
                  <Center>
                    <Grid><Text weight={700}>Name :  </Text> <Text> -  {enjoy[id].name}</Text></Grid>
                  </Center>
                  <Space h="xs" />
                  <Center>
                    <Grid><Text weight={700}>Rating :  </Text> <Text> -  {enjoy[id].rating}</Text></Grid>
                  </Center>
                  <Space h="xs" />
                  <Center>
                    <Grid><Text weight={700}>Vicinity :  </Text> <Text> -  {enjoy[id].vicinity}</Text></Grid>
                  </Center>
                  <Space h="xl" />
                </Paper>
              </>
            )
              : null
          }
        </div>

        <Space h="xl" />

        <Center>
          <Button onClick={goTravel} rightIcon={<ArrowBackUp size={18} />} compact variant="subtle" radius="xs">
            Go back
          </Button>
          <Button onClick={goSleep} data-testid="goSleep" rightIcon={<Bed size={18} />} compact variant="subtle" radius="xs">
            Search for Hotels
          </Button>
        </Center>
      </Card>

    </Container >


    // <Container size={720}>
    //   <Group grow spacing={0}>
    //     <Button variant="default" className={classes.button}>
    //       <ArrowForwardUp />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Car />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Train />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Walk />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Bike />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <PlaneInflight />
    //     </Button>
    //   </Group>

    //   <Space h="xl" />

    //   <AspectRatio ratio={16 / 9}>
    //     <iframe
    //       src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d10500.902039411167!2d2.2913514895690534!3d48.85391001859108!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x47e66e2964e34e2d%3A0x8ddca9ee380ef7e0!2sEiffel%20Tower!5e0!3m2!1sen!2sru!4v1653233639984!5m2!1sen!2sru"
    //       title="Google map"
    //       frameBorder="0"
    //     />
    //   </AspectRatio>
    // </Container>
  )
}

function enjoyActivities(enjoyActivities: any) {
  throw new Error("Function not implemented.");
}

