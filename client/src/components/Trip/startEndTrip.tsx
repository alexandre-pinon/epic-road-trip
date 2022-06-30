import { useNavigate } from 'react-router-dom';

import React, {useState} from "react";
import {ActionIcon, Button, Center, Container, createStyles, Modal, Space, TextInput, Title} from "@mantine/core";
import {BuildingSkyscraper, CalendarEvent, Search, BuildingCommunity, CalendarMinus} from "tabler-icons-react";
import {Calendar, DatePicker} from "@mantine/dates";
import {props} from "kea";



export function StartEndTrip() {

    //const [startCalendar, setStartCalendar] = useState(false);
    //const [endCalendar, setEndCalendar] = useState(false);
    const [startDateValue, setStartDate] = React.useState<Date | null>(new Date());
    const [endDateValue, setEndDate] = React.useState<Date | null>(new Date());
    const [startCity, setStartCity] = useState('Paris');
    const [endCity, setEndCity] = useState('');

    return (
        <>
            <Center>
                <Title order={2}> Start & End of your sub-trip </Title>
            </Center>

            <Space h="xl" />

            <Center>
                <Title order={3}> Please choose the city of departure and arrival of your sub-trip </Title>
            </Center>

            <Space h="lg" />


            <Container size={540}>
                {/* Ville de départ */}
                <TextInput
                    icon={<BuildingCommunity size={18} />}
                    radius="xl"
                    size="md"
                    placeholder="Ville de départ"
                    rightSectionWidth={42}
                    {...props}
                    value={startCity}
                    onChange={(event) => setStartCity(event.currentTarget.value)}
                />

                <Space h="xl" />

                {/* Ville d'arrivée */}
                <TextInput
                    icon={<BuildingSkyscraper size={18} />}
                    radius="xl"
                    size="md"
                    placeholder="Ville d'arrivée"
                    rightSectionWidth={42}
                    {...props}
                    onChange={(event) => setEndCity(event.currentTarget.value)}
                />

            </Container>
                <Space h="xl" />

                <Center>
                    <Title order={3}> Please choose the date of departure and arrival of your sub-trip </Title>
                </Center>

                <Space h="xl" />

            <Container size={540}>

                <DatePicker
                    icon={<CalendarEvent size={18} />}
                    placeholder="Start date of your Trip"
                    radius="xl"
                    onChange={setStartDate}
                />

                <Space h="xl" />

                <DatePicker
                    icon={<CalendarMinus size={18} />}
                    placeholder="End date of your Trip"
                    radius="xl"
                    onChange={setEndDate}
                />

                <Space h="xl" />

                <Center>
                    <Button rightIcon={<Search size={18} />} variant="light" radius="xl">
                        Search for itineraries
                    </Button>
                </Center>
            </Container>
        </>
    )
}