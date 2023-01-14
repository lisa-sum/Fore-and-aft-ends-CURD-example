'use client'
import { ReactNode } from 'react'
import Link from 'next/link'

import Login from './login/page'

import Create from './create/page'
import Read from './query/page'
import Update from './update/page'
import Delete from './delete/page'

type Routes = { path: string, label: string, element: ReactNode }[]

const routers: Routes = [
	{ path: 'login', label: 'Login', element: <Login /> },
	{ path: 'create', label: 'Create', element: <Create /> },
	{ path: 'read', label: 'Read', element: <Read /> },
	{ path: 'update', label: 'Update', element: <Update /> },
	{ path: 'delete', label: 'Delete', element: <Delete /> },
]

export default function Home () {
	return (
		<main>
			<ul>
				{
					routers.map((route) =>
						(<li key={ route.label }>
							<Link href={ `/${ route.path }` }>
							{ route.label }
						</Link>
						</li>),
					)
				}
			</ul>
		</main>
	)
}
