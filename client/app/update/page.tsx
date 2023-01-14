'use client'
import { Button } from '@mui/material'
import { useState } from 'react'

const updateOne = (
	key: string,
	value: string,
	update: string,
) => {
	console.log('key', key)
	console.log('value', value)
	console.log('update', update)

	if (key === '' || value === '' || update === ''){
		window.alert('更新参数都不能为空!')
		return
	}

	window.fetch('http://localhost:4000/user/update', {
		method: 'PATCH',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			key,
			value,
			update,
		}),
	})
	.then((res) => {

		if (!res.ok){
			const { statusText } = res
			throw new Error(statusText)
		}
		return res.json()
	})
	.then(res => {
		console.log(res)
	})
	.catch(err => {
		console.error(err)
	})
}

export default function Update () {
	const [key, setKey] = useState<string>('')
	const [value, setValue] = useState<string>('')
	const [update, setUpdate] = useState<string>('')
	return (
		<>
			<label htmlFor=''>Key:
				<input type='search' onChange={ (e) => setKey(e.currentTarget.value) } />
			</label>
			<label htmlFor=''>Value:
				<input type='search' onChange={ (e) => setValue(e.currentTarget.value) } />
			</label>
			<label htmlFor=''>Update Value:
				<input type='search' onChange={ (e) => setUpdate(e.currentTarget.value) } />
			</label>
			<Button onClick={ () => updateOne(key, value, update) }>Update</Button>
		</>
	)
};
