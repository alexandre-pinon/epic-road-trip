import { useNavigate } from 'react-router-dom';
import { Avatar, Button, Center, Title, Card, Image, Text, Paper, Space, Grid } from '@mantine/core';
import { useForm } from '@mantine/form';
import React, { useState } from "react";
import axios from "axios";





export function ResumeTrip({ fulTrip }: any) {
    const navigate = useNavigate();

    const confirmTripPost = () => {
        let params = [
            {
                city: fulTrip.startCity,
                startdate: fulTrip.startDateValue,
                enddate: fulTrip.endDateValue,
                enjoy: [
                    {
                        name: fulTrip.enjoyName.name,
                        rating: fulTrip.enjoyName.rating,
                        vicinity: fulTrip.enjoyName.vicinity
                    }
                ],
                sleep: [
                    {
                        name: fulTrip.sleep.nameSleep,
                        rating: fulTrip.sleep.ratingSleep,
                        vicinity: fulTrip.sleep.vicinitySleep
                    }
                ],
                eat: [
                    {
                        name: fulTrip.eat.name,
                        rating: fulTrip.eat.rating,
                        vicinity: fulTrip.eat.vicinity
                    }
                ],
                drink: [
                    {
                        name: fulTrip.drink.name,
                        rating: fulTrip.drink.rating,
                        vicinity: fulTrip.drink.vicinity
                    }
                ],
            },
            {
                city: fulTrip.endCity,
                startdate: fulTrip.startDateValue,
                enddate: fulTrip.endDateValue,
                travel: {
                    type: fulTrip.selectedTravel.type,
                    departure: {
                        city: fulTrip.selectedTravel.cityDeparture
                    },
                    arrival: {
                        city: fulTrip.selectedTravel.cityArrival
                    },
                    duration: fulTrip.selectedTravel.duration,
                    startdate: fulTrip.selectedTravel.startDate,
                    enddate: fulTrip.selectedTravel.endDate,
                },
                enjoy: [
                    {
                        name: fulTrip.enjoyArrival.name,
                        rating: fulTrip.enjoyArrival.rating,
                        vicinity: fulTrip.enjoyArrival.vicinity
                    }
                ],
                sleep: [
                    {
                        name: fulTrip.sleepArrival.nameSleep,
                        rating: fulTrip.sleepArrival.ratingSleep,
                        vicinity: fulTrip.sleepArrival.vicinitySleep
                    }
                ],
                eat: [
                    {
                        name: fulTrip.eatArrival.name,
                        rating: fulTrip.eatArrival.rating,
                        vicinity: fulTrip.eatArrival.vicinity
                    }
                ],
                drink: [
                    {
                        name: fulTrip.drinkArrival.name,
                        rating: fulTrip.drinkArrival.rating,
                        vicinity: fulTrip.drinkArrival.vicinity
                    }
                ],
            }
        ];
        axios({
            method: 'post',
            url: 'http://localhost:8000/api/v1/roadtrip/',
            params: {userID: '62c32834bdd39d78b24f4a70'},
            data: params
        }).then(res => {
            console.log(res)
        });
    };



    console.log(fulTrip)
    console.log(fulTrip.startDateValue.toDateString())



    const goToHome = async () => {
        console.log("Going back home!")
        navigate('/');
    };

    const goToArrival = async () => {
        navigate('/enjoyArrival');
    };


    const addSubTrip = async () => {
        console.log("Add a sub-trip")
        navigate('/');
    };

    return (
        <>
            <Center>
                <Title
                    align="center"
                    sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
                >
                    This is your resume trip 🚌

                </Title>
                <Space h="xl" />
            </Center>

            <Space h="xl" />

            <Paper shadow="xs" radius="lg" p="lg" withBorder>
                <Center><Title
                    order={2}
                    align="center"
                    sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
                >
                    Your trip

                </Title></Center>
                <h3>Cities</h3>
                <Grid><Text weight={700}>Start City :  </Text><Text> &nbsp;{fulTrip.startCity}</Text></Grid>
                <Grid><Text weight={700}>End City :  </Text><Text> &nbsp;{fulTrip.endCity}</Text></Grid>
                <Space h="md" />
                <h3>Dates</h3>
                <Grid><Text weight={700}>Start Date :  </Text><Text> &nbsp;{fulTrip.startDateValue.toDateString()}</Text></Grid>
                <Grid><Text weight={700}>End Date :  </Text><Text> &nbsp;{fulTrip.endDateValue.toDateString()}</Text></Grid>
                <Space h="md" />
                <h3>Itinerary</h3>
                <Grid><Text weight={700}>Departure city :  </Text><Text> &nbsp;{fulTrip.selectedTravel.cityDeparture}</Text></Grid>
                <Grid><Text weight={700}>Arrival city :  </Text><Text> &nbsp;{fulTrip.selectedTravel.cityArrival}</Text></Grid>
                <Grid><Text weight={700}>Departure time :  </Text><Text> &nbsp;{fulTrip.selectedTravel.startDate}</Text></Grid>
                <Grid><Text weight={700}>Arrival time :  </Text><Text> &nbsp;{fulTrip.selectedTravel.endDate}</Text></Grid>
                <Grid><Text weight={700}>Duration :  </Text><Text> &nbsp;{fulTrip.selectedTravel.duration}</Text></Grid>
                <Space h="md" />
                <h3>Enjoy Activities</h3>
                <Grid><Text weight={700}>Name :  </Text><Text> &nbsp;{fulTrip.enjoyName.name}</Text></Grid>
                <Grid><Text weight={700}>Vicinity :  </Text><Text> &nbsp;{fulTrip.enjoyName.vicinity}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text><Text> &nbsp;{fulTrip.enjoyName.rating}</Text></Grid>
                <Space h="md" />
                <h3>Eat</h3>
                <Grid><Text weight={700}>Restaurant name :  </Text><Text> &nbsp;{fulTrip.eat.name}</Text></Grid>
                <Grid><Text weight={700}>Address :  </Text><Text> &nbsp;{fulTrip.eat.vicinity}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text><Text> &nbsp;{fulTrip.eat.rating}</Text></Grid>
                <Space h="md" />
                <h3>Drink</h3>
                <Grid><Text weight={700}>Restaurant name :  </Text><Text> &nbsp;{fulTrip.drink.name}</Text></Grid>
                <Grid><Text weight={700}>Address :  </Text><Text> &nbsp;{fulTrip.drink.vicinity}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text><Text> &nbsp;{fulTrip.drink.rating}</Text></Grid>
                <Space h="md" />
                <h3>Hotel</h3>
                <Grid><Text weight={700}>Hotel :  </Text><Text> &nbsp;{fulTrip.sleep.nameSleep}</Text></Grid>
                <Grid><Text weight={700}>Address :  </Text><Text> &nbsp;{fulTrip.sleep.vicinitySleep}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text><Text> &nbsp;{fulTrip.sleep.ratingSleep}</Text></Grid>
                <Center><Button onClick={confirmTripPost}>Confirm your trip !</Button></Center>
            </Paper>

            {/*
            <Paper shadow="xs" radius="lg" p="lg" withBorder>
                {subTrip.map((subTrip, index) => (
                    <>
                        <Center><h2>Your sub-trip {index}</h2></Center>
                        <Paper shadow="md" radius="lg" p="lg" withBorder>
                            <Text>Start city : {subTrip.startCity}</Text>
                            <Text>End city : {subTrip.endCity}</Text>
                            <Text>Your mode of transport : {subTrip.Trajets}</Text>
                            <Text>{subTrip.Activities[0].hotel}</Text>
                            <Text>{subTrip.Activities[0].restaurant}</Text>
                            <Text>{subTrip.Activities[0].Entertainment1}</Text>
                            <Text>{subTrip.Activities[0].Entertainment1}</Text>

                            {subTrip.Activities.map((activities) => (
                                <>
                                    <Text>{activities.hotel}</Text>
                                    <Text>{activities.restaurant}</Text>
                                    <Text>{activities.Entertainment1}</Text>
                                    <Text>{activities.Entertainment2}</Text>
                                </>

                            ))}
                        </Paper>
                        <h1></h1>
                    </>
                ))}
            </Paper>
            */}


            <Space h="xl" />

            <Center>
                <Button onClick={goToHome} variant="light" radius="xl">
                    Go to Home Page
                </Button>
                <Button onClick={goToArrival} variant="light" radius="xl">
                    {fulTrip.endCity}
                </Button>
                {/*
                <Button onClick={addSubTrip} variant="light" radius="xl">
                    Add an other sub-Trip
                </Button>
                */}

            </Center>
        </>
    )
}