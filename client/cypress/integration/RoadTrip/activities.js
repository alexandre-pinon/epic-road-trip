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
})