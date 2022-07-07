import { useNavigate } from 'react-router-dom';
import { Avatar, Button, Center, Title, Card, Image, Text, Paper, Space, Grid, createStyles, Container, SimpleGrid, Divider } from '@mantine/core';
import { useForm } from '@mantine/form';
import React, { useState } from "react";
import axios from "axios";
import { ArrowBackUp, Building } from 'tabler-icons-react';


const useStyles = createStyles((theme) => ({
    wrapper: {
        paddingTop: 80,
        paddingBottom: 50,
    },

    item: {
        display: 'flex',
    },

    itemIcon: {
        padding: theme.spacing.xs,
        marginRight: theme.spacing.md,
    },

    itemTitle: {
        marginBottom: theme.spacing.xs / 2,
    },

    supTitle: {
        textAlign: 'center',
        textTransform: 'uppercase',
        fontWeight: 800,
        fontSize: theme.fontSizes.sm,
        color: theme.colors[theme.primaryColor][theme.colorScheme === 'dark' ? 4 : 8],
        letterSpacing: 0.5,
    },

    title: {
        lineHeight: 1,
        textAlign: 'center',
        marginTop: theme.spacing.xl,
    },

    description: {
        textAlign: 'center',
        marginTop: theme.spacing.xs,
    },

    highlight: {
        backgroundColor:
            theme.colorScheme === 'dark'
                ? theme.fn.rgba(theme.colors[theme.primaryColor][6], 0.55)
                : theme.colors[theme.primaryColor][0],
        padding: 5,
        paddingTop: 0,
        borderRadius: theme.radius.sm,
        display: 'inline-block',
        color: theme.colorScheme === 'dark' ? theme.white : 'inherit',
    },
}));




export function ResumeTrip({ fulTrip, auth }: any) {
    const navigate = useNavigate();
    const { classes } = useStyles();

    const [subTrip, setSubTrip] = useState([])

    const confirmTripPost = () => {
        console.log(auth)
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
            params: {userID: auth.userID},
            data: params
        }).then(res => {
            console.log(res)
        });
    };



    console.log(fulTrip)
    console.log(fulTrip.startDateValue.toDateString())



    const goToHome = async () => {
        console.log("Going back home!")
        navigate('/resumeTrip');
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
            <Container size={700} className={classes.wrapper}>
                <Center>
                    <Title
                        align="center"
                        data-testid="Big-title"
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
                        data-testid="yourTrip"
                        sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
                    >
                        Your trip


                    </Title></Center>
                    <Title order={5}>🌇</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Start City :  </Text><Text> &nbsp;{fulTrip.startCity}</Text></Grid>
                    <Grid><Text weight={300}>End City :  </Text><Text> &nbsp;{fulTrip.endCity}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>📅</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Start Date :  </Text><Text weight={400}> &nbsp;{fulTrip.startDateValue.toDateString()}</Text></Grid>
                    <Grid><Text weight={300}>End Date :  </Text><Text weight={400}> &nbsp;{fulTrip.endDateValue.toDateString()}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>✈️</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Departure city :  </Text><Text weight={400}> &nbsp;{fulTrip.selectedTravel.cityDeparture}</Text></Grid>
                    <Grid><Text weight={300}>Arrival city :  </Text><Text weight={400}> &nbsp;{fulTrip.selectedTravel.cityArrival}</Text></Grid>
                    <Grid><Text weight={300}>Departure time :  </Text><Text weight={400}> &nbsp;{fulTrip.selectedTravel.startDate}</Text></Grid>
                    <Grid><Text weight={300}>Arrival time :  </Text><Text weight={400}> &nbsp;{fulTrip.selectedTravel.endDate}</Text></Grid>
                    <Grid><Text weight={300}>Duration :  </Text><Text weight={400}> &nbsp;{fulTrip.selectedTravel.duration}</Text></Grid>
                    <Space h="xl" />
                    <Divider my="sm" />
                    <Center><Title
                        order={2}
                        align="center"
                        sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
                    >
                        {fulTrip.startCity}

                    </Title></Center>
                    <Space h="xl" />


                    <Title order={5}>🙂</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Name :  </Text><Text weight={400}> &nbsp;{fulTrip.enjoyName.name}</Text></Grid>
                    <Grid><Text weight={300}>Vicinity :  </Text><Text weight={400}> &nbsp;{fulTrip.enjoyName.vicinity}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.enjoyName.rating}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>🍔</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Restaurant name :  </Text><Text weight={400}> &nbsp;{fulTrip.eat.name}</Text></Grid>
                    <Grid><Text weight={300}>Address :  </Text><Text weight={400}> &nbsp;{fulTrip.eat.vicinity}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.eat.rating}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>🍺</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Restaurant name :  </Text><Text weight={400}> &nbsp;{fulTrip.drink.name}</Text></Grid>
                    <Grid><Text weight={300}>Address :  </Text><Text weight={400}> &nbsp;{fulTrip.drink.vicinity}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.drink.rating}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>😴</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Hotel :  </Text><Text weight={400}> &nbsp;{fulTrip.sleep.nameSleep}</Text></Grid>
                    <Grid><Text weight={300}>Address :  </Text><Text weight={400}> &nbsp;{fulTrip.sleep.vicinitySleep}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.sleep.ratingSleep}</Text></Grid>




                    <Space h="xl" />
                    <Divider my="sm" />
                    <Center><Title
                        order={2}
                        align="center"
                        sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 100 })}
                    >
                        {fulTrip.endCity}

                    </Title></Center>
                    <Space h="xl" />


                    {/* Arrival City */}
                    <Title order={5}>🙂</Title>
                    <Space h="xs" />

                    <Grid><Text weight={300}>Name :  </Text><Text weight={400}> &nbsp;{fulTrip.enjoyArrival.name}</Text></Grid>
                    <Grid><Text weight={300}>Vicinity :  </Text><Text weight={400}> &nbsp;{fulTrip.enjoyArrival.vicinity}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.enjoyArrival.rating}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>🍔</Title>
                    <Space h="xs" />

                    <Grid><Text weight={300}>Restaurant name :  </Text><Text weight={400}> &nbsp;{fulTrip.eatArrival.name}</Text></Grid>
                    <Grid><Text weight={300}>Address :  </Text><Text weight={400}> &nbsp;{fulTrip.eatArrival.vicinity}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.eatArrival.rating}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>🍺</Title>
                    <Space h="xs" />

                    <Grid><Text weight={300}>Restaurant name :  </Text><Text weight={400}> &nbsp;{fulTrip.drinkArrival.name}</Text></Grid>
                    <Grid><Text weight={300}>Address :  </Text><Text weight={400}> &nbsp;{fulTrip.drinkArrival.vicinity}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.drinkArrival.rating}</Text></Grid>
                    <Space h="md" />
                    <Title order={5}>😴</Title>
                    <Space h="xs" />
                    <Grid><Text weight={300}>Hotel :  </Text><Text weight={400}> &nbsp;{fulTrip.sleepArrival.nameSleep}</Text></Grid>
                    <Grid><Text weight={300}>Address :  </Text><Text weight={400}> &nbsp;{fulTrip.sleepArrival.vicinitySleep}</Text></Grid>
                    <Grid><Text weight={300}>Rating :  </Text><Text weight={400}> &nbsp;{fulTrip.sleepArrival.ratingSleep}</Text></Grid>
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
                    <Button onClick={goToHome} rightIcon={<ArrowBackUp size={18} />} compact variant="subtle" radius="xs">
                        Go to Home Page
                    </Button>
                    <Button onClick={confirmTripPost} compact variant="subtle" radius="xs">
                        Confirm your trip 🤩 !
                    </Button>
                    <Button onClick={goToArrival} rightIcon={<Building size={18} />} compact variant="subtle" radius="xs">
                        {fulTrip.endCity}
                    </Button>
                    {/*
                <Button onClick={addSubTrip} variant="light" radius="xl">
                    Add an other sub-Trip
                </Button>
                */}

                </Center>
            </Container>

        </>
    )
}