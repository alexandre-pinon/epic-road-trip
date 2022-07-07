import {
  AspectRatio,
  Button,
  Card,
  Container,
  createStyles,
  Group,
  Space,
  Tooltip,
  Text,
  SimpleGrid,
  Paper,
  Grid, Center, ActionIcon, Title
} from "@mantine/core";
import { ArrowBackUp, ArrowForwardUp, Bike, Car, MoodSmile, PlaneInflight, Search, Train, Trash, Walk } from 'tabler-icons-react';
import axios from "axios";
import React, { SetStateAction, useEffect, useState } from "react";
import { use } from "msw/lib/types/utils/internal/requestHandlerUtils";
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
}));


export function Travel({ fulTrip }: any) {
  const { classes } = useStyles();
  const navigate = useNavigate();

  const goStartAndEnd = async () => {
    console.log("Go to travel page!")
    console.log(fulTrip)
    console.log(fulTrip.startDateValue)
    //navigate('/startEndTrip');
  };

  const goEnjoy = async () => {
    console.log("Go to enjoy page!")
    navigate('/enjoy');
  };

  const goResumeTrip = async () => {
    console.log("Go to enjoy page!")
    navigate('/resumeTrip');
  };

  const [id, setId] = useState(0)
  const [cityDeparture, setCityDeparture] = useState('')
  const [cityArrival, setCityArrival] = useState('')
  const [duration, setDuration] = useState('')
  const [startDate, setStartDate] = useState('')
  const [endDate, setEndDate] = useState('')
  const [price, setPrice] = useState('')

  let [selectedTravel, setSelectedTravel] = useState('')

  const [togglePlane, setTogglePlane] = useState(false)
  const [toggleTrain, setToggleTrain] = useState(false)



  const [plane, setPlane] = useState([{
    type: "AIR",
    id,
    cityDeparture,
    cityArrival,
    duration,
    startDate,
    endDate
  }
  ])

  const [train, setTrain] = useState([{
    type: "GROUND",
    id,
    cityDeparture,
    cityArrival,
    duration,
    startDate,
    endDate
  }
  ])




  const planeTravel = () => {
    let params = {
      adults: 1,
      departureDate: fulTrip.startDateValue,
      destinationLocation: fulTrip.endCity,
      maxPrice: 10000,
      originLocation: fulTrip.startCity
    };
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/travel/air',
      data: params
    }).then(res => {
      let id = 0
      let myList2: any = [];
      res.data.data.forEach((data: any) => {
        let travelInfo = {
          id: id,
          cityDeparture: data.departure.city,
          cityArrival: data.arrival.city,
          duration: data.duration,
          startDate: data.startdate,
          endDate: data.enddate
        }
        myList2.push(travelInfo)
        id++
      })
      setPlane(myList2)
      console.log(myList2)
    });
  };

  const trainTravel = () => {
    let params = {
      departureTime: fulTrip.startDateValue,
      destination: fulTrip.endCity,
      origin: fulTrip.startCity
    };
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/travel/ground',
      data: params
    }).then(res => {
      let id = 0
      let myList2: any = [];
      res.data.data.forEach((data: any) => {
        let travelInfo = {
          id: id,
          cityDeparture: data.departure.city,
          cityArrival: data.arrival.city,
          duration: data.duration,
          startDate: data.startdate,
          endDate: data.enddate,
        }
        myList2.push(travelInfo)
        id++;
      })
      setTrain(myList2)
      console.log(myList2)
    });
  };

  useEffect(() => {
    planeTravel();
    trainTravel();
  }, []);


  const showPlane = () => {
    setTogglePlane(!togglePlane)
  }

  const showTrain = () => {
    setToggleTrain(!toggleTrain)
  }

  const test = () => {
    console.log("test")
  }

  const SelectTravel = (id: number, type: string) => {

    setSelectedTravel(type)
    console.log(id)
    setId(id)

    if (type == "Plane") {
      fulTrip.setSelectedTravel(plane[id])
    }
    else {
      fulTrip.setSelectedTravel(train[id])
    }
    console.log("scroll to bottom: ", document.body.offsetHeight)
    window.scroll({
      top: document.body.offsetHeight,
      left: 0,
      behavior: 'smooth',
    });
  }

  return (
    <Container size={600}>
      <Group grow spacing={0} position="apart">

        <SimpleGrid cols={1} spacing="sm" breakpoints={[{ maxWidth: 'sm', cols: 1 }]}>
          <>
            <Button onClick={showPlane} variant="default" className={classes.button}>
              <PlaneInflight />
            </Button>

            <ul>
              {
                togglePlane ? (
                  plane.map((item) => (
                    <Paper shadow="xl" p="md" withBorder key={item.id}>
                      <Grid><Text weight={500}>Departure :  </Text> <Text weight={400}>&nbsp;{item.cityDeparture}</Text></Grid>                      <Space h="md" />
                      <Grid><Text weight={500}>Arrival :  </Text> <Text weight={400}> &nbsp;{item.cityArrival}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={500}>Start Date :  </Text> <Text weight={400}>&nbsp;{item.startDate}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={500}>End Date :  </Text> <Text weight={400}>&nbsp;{item.endDate}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={500}>Duration :  </Text> <Text weight={400}>&nbsp;{item.duration}</Text></Grid>
                      <Space h="xl" />

                      <Center>
                        <ActionIcon onClick={() => SelectTravel(item.id, 'Plane')} variant="outline">👆</ActionIcon>

                      </Center>
                    </Paper>
                  ))
                )
                  : null
              }
            </ul>
          </>
        </SimpleGrid>

      </Group>

      <Space h="sm" />

      <Group grow spacing={0} position="apart">
        <SimpleGrid cols={1} spacing="sm" breakpoints={[{ maxWidth: 'sm', cols: 1 }]}>
          <>
            <Button onClick={showTrain} variant="default" className={classes.button}>
              <Train />
            </Button>

            <ul>
              {
                toggleTrain ? (
                  train.map((item) => (
                    <Paper shadow="xl" p="md" withBorder key={item.id}>
                      <Grid><Text weight={700}>Departure :  </Text> <Text>&nbsp;{item.cityDeparture}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={700}>Arrival :  </Text> <Text>&nbsp;{item.cityArrival}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={700}>Start Date :  </Text> <Text>&nbsp;{item.startDate}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={700}>End Date :  </Text> <Text>&nbsp;{item.endDate}</Text></Grid>
                      <Space h="md" />

                      <Grid><Text weight={700}>Duration :  </Text> <Text>&nbsp;{item.duration}</Text></Grid>
                      <Space h="xl" />
                      <Center>
                        <ActionIcon onClick={() => SelectTravel(item.id, 'Train')} variant="outline">👆</ActionIcon>

                      </Center>
                    </Paper>
                  ))
                )
                  : null
              }
            </ul>
          </>
        </SimpleGrid>
      </Group>


      <div>
        {
          selectedTravel === "Plane" ? (
            <>
              <Space h="xs" />

              <Title
                order={2}
                align="center"
                sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
              >
                Selected Travel


              </Title>
              <Paper shadow="xl" p="md" withBorder >
                <Grid><Text weight={700}>Travel : PLANE number  {plane[id].id}</Text> </Grid>
                <Space h="xs" />

                <Grid><Text weight={700}>Departure :  </Text> <Text>&nbsp;{plane[id].cityDeparture}</Text></Grid>
                <Space h="xs" />

                <Grid><Text weight={700}>Arrival :  </Text> <Text>&nbsp;{plane[id].cityArrival}</Text></Grid>
                <Space h="xs" />

                <Grid><Text weight={700}>Start Date :  </Text> <Text>&nbsp;{plane[id].startDate}</Text></Grid>
                <Space h="xs" />

                <Grid><Text weight={700}>End Date :  </Text> <Text>&nbsp;{plane[id].endDate}</Text></Grid>
                <Space h="xs" />

                <Grid><Text weight={700}>Duration :  </Text> <Text>&nbsp;{plane[id].duration}</Text></Grid>
              </Paper>
            </>
          )
            : null
        }
        {
          selectedTravel === "Train" ? (
            <>
              <Space h="xs" />

              <Title
                order={2}
                align="center"
                sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
              >
                Selected Travel


              </Title>
              <Paper shadow="xl" p="md" withBorder >
                <Grid><Text weight={700}>Travel : TRAIN number  {train[id].id}</Text> </Grid>
                <Grid><Text weight={700}>Departure :  </Text> <Text>&nbsp;{train[id].cityDeparture}</Text></Grid>
                <Grid><Text weight={700}>Arrival :  </Text> <Text>&nbsp;{train[id].cityArrival}</Text></Grid>
                <Grid><Text weight={700}>Start Date :  </Text> <Text>&nbsp;{train[id].startDate}</Text></Grid>
                <Grid><Text weight={700}>End Date :  </Text> <Text>&nbsp;{train[id].endDate}</Text></Grid>
                <Grid><Text weight={700}>Duration :  </Text> <Text>&nbsp;{train[id].duration}</Text></Grid>
              </Paper>
            </>

          )
            : null
        }
      </div>



      <Space h="xl" />


      <Center>
        <Button onClick={goStartAndEnd} rightIcon={<ArrowBackUp size={18} />} compact variant="subtle" radius="xs">
          Go back
        </Button>
        <Button onClick={goEnjoy} rightIcon={<MoodSmile size={18} />} compact variant="subtle" radius="xs">
          Search for Activities
        </Button>
        {/* <Button onClick={goResumeTrip} rightIcon={<Trash size={18} />} compact variant="subtle" radius="xs">
          Delete après
        </Button> */}
      </Center>


      {/*
      <AspectRatio ratio={16 / 9}>
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d10500.902039411167!2d2.2913514895690534!3d48.85391001859108!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x47e66e2964e34e2d%3A0x8ddca9ee380ef7e0!2sEiffel%20Tower!5e0!3m2!1sen!2sru!4v1653233639984!5m2!1sen!2sru"
          title="Google map"
          frameBorder="0"
        />
      </AspectRatio>
      */}


    </Container>
  )
}