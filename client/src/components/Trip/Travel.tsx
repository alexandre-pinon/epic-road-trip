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
  Grid, Center
} from "@mantine/core";
import {ArrowForwardUp, Bike, Car, PlaneInflight, Search, Train, Walk} from 'tabler-icons-react';
import axios from "axios";
import React, {SetStateAction, useEffect, useState} from "react";
import {use} from "msw/lib/types/utils/internal/requestHandlerUtils";
// @ts-ignore
import { v4 as uuidv4 } from 'uuid';
import {useNavigate} from "react-router-dom";

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


export function Travel({fulTrip} : any) {
  const {classes} = useStyles();
  const navigate = useNavigate();

  const goStartAndEnd = async () => {
    console.log("Go to travel page!")
    navigate('/startEndTrip');
  };

  const goEnjoy = async () => {
    console.log("Go to enjoy page!")
    navigate('/enjoy');
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
    id,
    cityDeparture,
    cityArrival,
    duration,
    startDate,
    endDate
  }
  ])

  const [train, setTrain] = useState([{
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
      departureDate: "2022-08-08T15:04:05Z",
      destinationLocation: "Paris",
      maxPrice: 10000,
      originLocation: "London",
    };
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/travel/air',
      data: params
    }).then(res => {
      let id = 0
      let myList2: any = [];
      res.data.data.forEach((data: any ) => {
        let travelInfo = {
          id: id,
          cityDeparture: data.arrival.city,
          cityArrival: data.departure.city,
          duration: data.duration,
          startDate: data.startdate,
          endDate: data.enddate
        }
        myList2.push(travelInfo)
        id ++
      })
      setPlane(myList2)
      console.log(myList2)
    });
  };

  const trainTravel = () => {
    let params = {
      departureTime: "2022-08-18T15:04:05Z",
      destination: "Paris",
      origin: "London",
    };
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/travel/ground',
      data: params
    }).then(res => {
      let id = 0
      let myList2: any = [];
      res.data.data.forEach((data: any ) => {
        let travelInfo = {
          id: id,
          cityDeparture: data.arrival.city,
          cityArrival: data.departure.city,
          duration: data.duration,
          startDate: data.startdate,
          endDate: data.enddate,
        }
        myList2.push(travelInfo)
        id ++;
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
  }

  return (
      <Container size={600}>
        <Group grow spacing={0} position="apart">
            <SimpleGrid cols={1}>
              <>
                <Button onClick={showPlane} variant="default" className={classes.button}>
                  <PlaneInflight />
                </Button>

                <ul>
                  {
                    togglePlane ? (
                            plane.map((item) => (
                                  <Paper shadow="xl" p="md" withBorder key={item.id}>
                                    <Grid><Text weight={700}>Travel : PLANE number  {item.id}</Text> </Grid>
                                    <Grid><Text weight={700}>Departure :  </Text> <Text> -  {item.cityDeparture}</Text></Grid>
                                    <Grid><Text weight={700}>Arrival :  </Text> <Text> -  {item.cityArrival}</Text></Grid>
                                    <Grid><Text weight={700}>Start Date :  </Text> <Text> -  {item.startDate}</Text></Grid>
                                    <Grid><Text weight={700}>End Date :  </Text> <Text> -  {item.endDate}</Text></Grid>
                                    <Grid><Text weight={700}>Duration :  </Text> <Text> -  {item.duration}</Text></Grid>
                                    <Center><Button onClick={()=>SelectTravel(item.id, 'Plane')} >Select this travel </Button></Center>
                                  </Paper>
                            ))
                        )
                        : null
                  }
                </ul>
              </>
            </SimpleGrid>

          </Group>

          <Group grow spacing={0} position="apart">
            <SimpleGrid cols={1}>
              <>
                <Button onClick={showTrain} variant="default" className={classes.button}>
                  <Train />
                </Button>
                <ul>
                  {
                    toggleTrain ? (
                            train.map((item) => (
                                <Paper shadow="xl" p="md" withBorder key={item.id}>
                                  <Grid><Text weight={700}>Travel : TRAIN number {item.id} </Text> </Grid>
                                  <Grid><Text weight={700}>Departure :  </Text> <Text> -  {item.cityDeparture}</Text></Grid>
                                  <Grid><Text weight={700}>Arrival :  </Text> <Text> -  {item.cityArrival}</Text></Grid>
                                  <Grid><Text weight={700}>Start Date :  </Text> <Text> -  {item.startDate}</Text></Grid>
                                  <Grid><Text weight={700}>End Date :  </Text> <Text> -  {item.endDate}</Text></Grid>
                                  <Grid><Text weight={700}>Duration :  </Text> <Text> -  {item.duration}</Text></Grid>
                                  <Center><Button onClick={()=>SelectTravel(item.id, 'Train')} >Select this travel </Button></Center>
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
            selectedTravel==="Plane" ? (
                <>
                  <Center><h3>SELECTED TRAVEL : </h3></Center>
                  <Paper shadow="xl" p="md" withBorder >
                    <Grid><Text weight={700}>Travel : PLANE number  {plane[id].id}</Text> </Grid>
                    <Grid><Text weight={700}>Departure :  </Text> <Text> -  {plane[id].cityDeparture}</Text></Grid>
                    <Grid><Text weight={700}>Arrival :  </Text> <Text> -  {plane[id].cityArrival}</Text></Grid>
                    <Grid><Text weight={700}>Start Date :  </Text> <Text> -  {plane[id].startDate}</Text></Grid>
                    <Grid><Text weight={700}>End Date :  </Text> <Text> -  {plane[id].endDate}</Text></Grid>
                    <Grid><Text weight={700}>Duration :  </Text> <Text> -  {plane[id].duration}</Text></Grid>
                  </Paper>
                </>
                )
                : null
          }
          {
            selectedTravel==="Train" ? (
                <>
                  <Center><h3>SELECTED TRAVEL : </h3></Center>
                  <Paper shadow="xl" p="md" withBorder >
                  <Grid><Text weight={700}>Travel : TRAIN number  {train[id].id}</Text> </Grid>
                  <Grid><Text weight={700}>Departure :  </Text> <Text> -  {train[id].cityDeparture}</Text></Grid>
                  <Grid><Text weight={700}>Arrival :  </Text> <Text> -  {train[id].cityArrival}</Text></Grid>
                  <Grid><Text weight={700}>Start Date :  </Text> <Text> -  {train[id].startDate}</Text></Grid>
                  <Grid><Text weight={700}>End Date :  </Text> <Text> -  {train[id].endDate}</Text></Grid>
                  <Grid><Text weight={700}>Duration :  </Text> <Text> -  {train[id].duration}</Text></Grid>
                </Paper>
                </>

                )
                : null
          }
        </div>







        <Space h="xl" />


        <Center>
          <Button onClick={goStartAndEnd} rightIcon={<Search size={18} />} variant="light" radius="xl">
            Go back
          </Button>
          <Button onClick={goEnjoy} rightIcon={<Search size={18} />} variant="light" radius="xl">
            Search for Activities
          </Button>
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