import React, { useState } from 'react'
import { kea, actions, path, reducers, useActions, useValues, listeners, afterMount, selectors, props } from 'kea';
import { createStyles, Container, UnstyledButton, Text, Card, SimpleGrid, Overlay, Space, Input, Center, TextInput, ActionIcon, Modal, Button, Title, Loader, List, ThemeIcon } from '@mantine/core';
import { Calendar } from '@mantine/dates';

import {
  HotelService,
  BrandAirbnb,
  Run,
  BrandTripadvisor,
  BrandBooking,
  Plus,
  Search,
  GasStation,
  Gauge,
  ManualGearbox,
  Users,
  CalendarEvent,
  BuildingSkyscraper,
  UserSearch,
  ArrowNarrowRight,
} from 'tabler-icons-react';

import type { logicType } from "./HomeType";
import {useNavigate} from "react-router-dom";

const API_URL = 'https://api.github.com'

const logic = kea<logicType>([
  path(["src\\components\\Home\\Home\\Home"]),
  actions({
    setUsername: (username) => ({ username }),
    setRepositories: (repositories) => ({ repositories }),
    setFetchError: (error) => ({ error }),
  }),

  reducers({
    username: [
      'keajs',
      {
        setUsername: (_, { username }) => username,
      },
    ],
    repositories: [
      [],
      {
        setUsername: () => [],
        setRepositories: (_, { repositories }) => repositories,
      },
    ],
    isLoading: [
      false,
      {
        setUsername: () => true,
        setRepositories: () => false,
        setFetchError: () => false,
      },
    ],
    error: [
      null,
      {
        setUsername: () => null,
        setFetchError: (_, { error }) => error,
      },
    ],
  }),

  selectors({
    sortedRepositories: [
      (selectors) => [selectors.repositories],
      (repositories) => {
        return [...repositories].sort((a, b) => b.stargazers_count - a.stargazers_count)
      },
    ],
  }),

  listeners(({ actions }) => ({
    setUsername: async ({ username }, breakpoint) => {
      await breakpoint(300)

      const url = `${API_URL}/users/${username}/repos?per_page=250`

      let response
      try {
        response = await window.fetch(url)
      } catch (error) {
        actions.setFetchError(error)
        return
      }

      breakpoint()

      const json = await response.json()

      if (response.status === 200) {
        actions.setRepositories(json)
      } else {
        actions.setFetchError(json.message)
      }
    },
  })),

  afterMount(({ actions, values }) => {
    actions.setUsername(values.username)
  }),
])

const mockdata = [
  { title: 'Hotels', icon: HotelService, color: 'dark' },
  { title: 'Vacation Rentals', icon: BrandAirbnb, color: 'dark' },
  { title: 'Things To Do', icon: Run, color: 'dark' },
  { title: 'Restaurants', icon: BrandTripadvisor, color: 'dark' },
  { title: 'Travel Forums', icon: BrandBooking, color: 'dark' },
  { title: 'More ...', icon: Plus, color: 'dark' },
];

const mockdata2 = [
  { label: '4 passengers', icon: Users },
  { label: '100 km/h in 4 seconds', icon: Gauge },
  { label: 'Automatic gearbox', icon: ManualGearbox },
  { label: 'Electric', icon: GasStation },
];

const useStyles = createStyles((theme) => ({
  card: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
  },

  title: {
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
    fontWeight: 700,
  },

  item: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: 'center',
    borderRadius: theme.radius.md,
    height: 90,
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
    transition: 'box-shadow 150ms ease, transform 100ms ease',

    '&:hover': {
      boxShadow: `${theme.shadows.md} !important`,
      transform: 'scale(1.05)',
    },
  },

  // [START] Hero with Background Image
  wrapper: {
    position: 'relative',
    paddingTop: 180,
    paddingBottom: 130,
    backgroundImage:
      'url(https://images2.alphacoders.com/106/1062364.png)',
    backgroundSize: 'cover',
    backgroundPosition: 'center',

    '@media (max-width: 520px)': {
      paddingTop: 80,
      paddingBottom: 50,
    },
  },

  inner: {
    position: 'relative',
    zIndex: 1,
  },
  // [END] Hero with Background Image

  // [START] Card with icon features
  card2: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
  },

  imageSection: {
    padding: theme.spacing.md,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    borderBottom: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
      }`,
  },

  label: {
    marginBottom: theme.spacing.xs,
    lineHeight: 1,
    fontWeight: 700,
    fontSize: theme.fontSizes.xs,
    letterSpacing: -0.25,
    textTransform: 'uppercase',
  },

  section: {
    padding: theme.spacing.md,
    borderTop: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
      }`,
  },

  icon: {
    marginRight: 5,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.gray[5],
  },
  // [END] Card with icon features
}));

export function Home({fulTrip} : any) {
  const navigate = useNavigate();

  // const [username, setUsername] = useState('keajs')
  const { username, isLoading, sortedRepositories, error } = useValues(logic)
  const { setUsername } = useActions(logic)
  const { classes, theme, cx } = useStyles();
  const [startCalendar, setStartCalendar] = useState(false);
  const [endCalendar, setEndCalendar] = useState(false);
  const [startDateValue, setStartDate] = React.useState<Date | null>(new Date());
  const [endDateValue, setEndDate] = React.useState<Date | null>(new Date());


  const validStartCity = (e: any) => {

    console.log(e.key)
    if(e.key === "Enter") {
      fulTrip.setStartCity(e.target.value)
      navigate('/startEndTrip');

    }
  }

  const items = mockdata.map((item) => (
    <UnstyledButton key={item.title} className={classes.item}>
      <item.icon color={theme.colors[item.color][6]} size={32} />
      <Text size="xs" mt={7}>
        {item.title}
      </Text>
    </UnstyledButton>
  ));

  const features = mockdata2.map((feature) => (
    <Center key={feature.label}>
      <feature.icon size={18} className={classes.icon} />
      <Text size="xs">{feature.label}</Text>
    </Center>
  ));

  return (
    <Container>
      <Card withBorder radius="md" className={classes.card}>
        <SimpleGrid cols={6} mt="md">
          {items}
        </SimpleGrid>
      </Card>

      <Space h="xl" />

      <div className={classes.wrapper}>
        <Overlay color="#000" opacity={0.25} zIndex={1} />

        <div className={classes.inner}>
          <Container size={640}>
            <Input
              icon={<Search />}
              placeholder="Where to?"
              radius="xl"
              size="md"
              onKeyPress={(e: any) => validStartCity(e)}
            />
          </Container>
        </div>
      </div>

      <Space h="xl" />
      <Space h="xl" />
      <Space h="xl" />

      <Container size={540}>
        {/* Ville de départ */}
        <TextInput
          icon={<BuildingSkyscraper size={18} />}
          radius="xl"
          size="md"
          rightSection={
            <ActionIcon size={32} radius="xl" color={theme.primaryColor} variant="hover">
              <Modal
                centered
                overlayColor={theme.colorScheme === 'dark' ? theme.colors.dark[9] : theme.colors.gray[2]}
                overlayOpacity={0.55}
                overlayBlur={3}
                opened={startCalendar}
                onClose={() => setStartCalendar(false)}
                withCloseButton={false}
              >
                <Calendar value={startDateValue} onChange={setStartDate} />;
              </Modal>
              <CalendarEvent size={18} onClick={() => setStartCalendar(true)} />
            </ActionIcon>
          }
          placeholder="Ville de départ"
          rightSectionWidth={42}
          {...props}
        />

        <Space h="xl" />

        {/* Ville d'arrivée */}
        <TextInput
          icon={<BuildingSkyscraper size={18} />}
          radius="xl"
          size="md"
          rightSection={
            <ActionIcon size={32} radius="xl" color={theme.primaryColor} variant="hover">
              <Modal
                centered
                overlayColor={theme.colorScheme === 'dark' ? theme.colors.dark[9] : theme.colors.gray[2]}
                overlayOpacity={0.55}
                overlayBlur={3}
                opened={endCalendar}
                onClose={() => setEndCalendar(false)}
                withCloseButton={false}
              >
                <Calendar value={endDateValue} onChange={setEndDate} />;
              </Modal>
              <CalendarEvent size={18} onClick={() => setEndCalendar(true)} />
            </ActionIcon>
          }
          placeholder="Ville d'arrivée"
          rightSectionWidth={42}
          {...props}
        />

        <Space h="xl" />

        <Center>
          <Button rightIcon={<Search size={18} />} variant="light" radius="xl">
            Rechercher
          </Button>
        </Center>

      </Container>

      {/* <Title order={3}>Top experiences on Epic Road Trip</Title> */}

      {/* GitHub API Test*/}
      <div>
        <div>
          <Title style={{ color: "#616161 " }} order={4}>Search for a GitHub user</Title>
          <Input
            icon={<UserSearch size={16} />}
            placeholder="Search user"
            value={username}
            type="text"
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setUsername(e.target.value)}
          />
        </div>

        {isLoading ? (
          <Loader size="sm" variant="dots" />
        ) : sortedRepositories.length > 0 ? (
          <Text color="dimmed">
            Found {sortedRepositories.length} repositories for user {username}!
            {sortedRepositories.map((repo) => (
              <List
                spacing="xs"
                size="sm"
                center
                icon={
                  <ThemeIcon color="gray" size={18} radius="xl">
                    <ArrowNarrowRight size={16} />
                  </ThemeIcon>
                }
                key={repo.id}>
                <Space h="xs" />
                <List.Item>
                  <a href={repo.html_url} target="_blank" rel="noreferrer">
                    {repo.full_name}
                  </a>
                  {' - '}
                  {repo.stargazers_count} stars, {repo.forks} forks.
                </List.Item>
              </List>
            ))}
          </Text>
        ) : (
          <div>{error ? `Error: ${error}` : 'No repositories found'}</div>
        )}
      </div>
    </Container>
  );
}

