describe('Activities pages', () => {

    it('Enjoy Activities', () => {
        cy.visit('http://localhost:3000/enjoy')

        cy.get('[data-testid="title"]').contains("Enjoy Activities")

        cy.get('[data-testid="goSleep"]').click()
        cy.location('href').should('include', '/sleep')
    })



    it('Sleep Activities', () => {
        cy.visit('http://localhost:3000/sleep')

        cy.get('[data-testid="title"]').contains("Sleep Activities")

        cy.get('[data-testid="goEat"]').click()
        cy.location('href').should('include', '/eat')

        cy.visit('http://localhost:3000/sleep')
        cy.get('[data-testid="goBack"]').click()
        cy.location('href').should('include', '/enjoy')
    })



    it('Eat activities', () => {
        cy.visit('http://localhost:3000/eat')

        cy.get('[data-testid="title"]').contains("Eat Activities")

        cy.get('[data-testid="goDrink"]').click()
        cy.location('href').should('include', '/drink')

        cy.visit('http://localhost:3000/eat')
        cy.get('[data-testid="goBack"]').click()
        cy.location('href').should('include', '/sleep')
    })



    it('Drink activities', () => {
        cy.visit('http://localhost:3000/drink')

        cy.get('[data-testid="title"]').contains("Drink Activities")

        cy.get('[data-testid="enjoyArrival"]').click()
        cy.location('href').should('include', '/enjoyArrival')

        cy.visit('http://localhost:3000/drink')
        cy.get('[data-testid="goBack"]').click()
        cy.location('href').should('include', '/eat')
    })



    it('enjoyArrival activities', () => {
        cy.visit('http://localhost:3000/enjoyArrival')

        cy.get('[data-testid="title"]').contains("Enjoy Activities")

        cy.get('[data-testid="sleepArrival"]').click()
        cy.location('href').should('include', '/sleepArrival')
    })




    it('sleepArrival activities', () => {
        cy.visit('http://localhost:3000/sleepArrival')

        cy.get('[data-testid="title"]').contains("Sleep Activities")

        cy.get('[data-testid="goEat"]').click()
        cy.location('href').should('include', '/eatArrival')

        cy.visit('http://localhost:3000/sleepArrival')
        cy.get('[data-testid="goBack"]').click()
        cy.location('href').should('include', '/enjoyArrival')
    })



    it('eatArrival activities', () => {
        cy.visit('http://localhost:3000/eatArrival')

        cy.get('[data-testid="title"]').contains("Eat Activities")

        cy.get('[data-testid="drinkArrival"]').click()
        cy.location('href').should('include', '/drinkArrival')

        cy.visit('http://localhost:3000/eatArrival')
        cy.get('[data-testid="goBack"]').click()
        cy.location('href').should('include', '/sleepArrival')
    })



    it('drinkArrival activities', () => {
        cy.visit('http://localhost:3000/drinkArrival')

        cy.get('[data-testid="title"]').contains("Drink Activities")

        cy.get('[data-testid="ResumeTrip"]').click()
        cy.location('href').should('include', '/resumeTrip')

        cy.visit('http://localhost:3000/drinkArrival')
        cy.get('[data-testid="goBack"]').click()
        cy.location('href').should('include', '/eatArrival')
    })
})