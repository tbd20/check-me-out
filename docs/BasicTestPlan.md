# Basic Test Plan
This document will try and outline the basic steps I am going to use to write and test the checkout service I will implement. 

## Main Goals
The main goal is to satisfy the requirements laid out in the problem statement. Allowing you to add items to a checkout basket and retrieve the total for said basket. My testing will be split between these two functions. 

1. Scan
For the scan function we need to ensure that multiple Items can be added to the basket and in any order in any order.

Some test cases that I can think of are listed below:
- As per the problem statement each element in the list is represented by a single letter. As per this spec I will enforce this as a rule. 
- Upper and lowercase input will be ignored. This means that Scan(A) and Scan(a) will be considered the same scan.
- Items can be added in any order.
- Scanning Items that are not in the store will return an error. 

2. GetTotalPrice
For the get total price function we will need to ensure that any combination of inputs in the Scan function lead to the correct total price. 

Some test cases are below: 
- Inputing one item will lead to correct result.
- Inputing two different items leads to the correct result.
- When inputing mulitple items, items scanned in any order lead to the correct result.
- A single item scanned enough times to reach the multi buy discount contributes the special price amount.
- Items beyond the special price amount are added at the unit price.
- Items beyond the special price amount are added at the multibuy amount if enough of the items are scanned. E.g. In the Problem statement Buying 3 of SKU A will lead to a price of 260.
- As we are told that the price changes frequently if there is a price change between the scanning of items and the GetTotalPrice function being invoked this change should be reflected.