import React, { useState } from 'react'
import { createStyles, Container, Image, Grid, UnstyledButton, Text, Card, Group, SimpleGrid, Button, Overlay, Title, Space, Input, Center } from '@mantine/core';

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
  Badge,
  UserSearch,
} from 'tabler-icons-react';

const mockdata = [
  { title: 'Hotels', icon: HotelService, color: 'dark' },
  { title: 'Vacation Rentals', icon: BrandAirbnb, color: 'dark' },
  { title: 'Things To Do', icon: Run, color: 'dark' },
  { title: 'Restaurants', icon: BrandTripadvisor, color: 'dark' },
  { title: 'Travel Forums', icon: BrandBooking, color: 'dark' },
  { title: 'More ...', icon: Plus, color: 'dark' },
  // { title: 'Reports', icon: Report, color: 'pink' },
  // { title: 'Payments', icon: Coin, color: 'red' },
  // { title: 'Cashback', icon: CashBanknote, color: 'orange' },
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

export function Home() {
  const [username, setUsername] = useState('keajs')
  const { classes, theme, cx } = useStyles();


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
            />
          </Container>
        </div>
      </div>

      <Space h="xl" />
      <Space h="xl" />
      <Space h="xl" />

      {/* <Title order={3}>Top experiences on Epic Road Trip</Title> */}

      {/* [START] Card with icon features */}
      {/* <Card withBorder radius="md" className={classes.card}>
        <Card.Section className={classes.imageSection}>
          <Image src="https://i.imgur.com/ZL52Q2D.png" alt="Tesla Model S" />
        </Card.Section>

        <Group position="apart" mt="md">
          <div>
            <Text weight={500}>Tesla Model S</Text>
            <Text size="xs" color="dimmed">
              Free recharge at any station
            </Text>
          </div>
          <Badge>25% off</Badge>
        </Group>

        <Card.Section className={classes.section} mt="md">
          <Text size="sm" color="dimmed" className={classes.label}>
            Basic configuration
          </Text>

          <Group spacing={8} mb={-8}>
            {features}
          </Group>
        </Card.Section>

        <Card.Section className={classes.section}>
          <Group spacing={30}>
            <div>
              <Text size="xl" weight={700} sx={{ lineHeight: 1 }}>
                $168.00
              </Text>
              <Text size="sm" color="dimmed" weight={500} sx={{ lineHeight: 1 }} mt={3}>
                per day
              </Text>
            </div>

            <Button radius="xl" style={{ flex: 1 }}>
              Rent now
            </Button>
          </Group>
        </Card.Section>
      </Card> */}
      {/* [END] Card with icon features*/}

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
        <Text color="dimmed">Repos will come here for user <strong>{username}</strong></Text>
      </div>
    </Container>
  );
}

