import Head from 'next/head'
import Image from 'next/image'
import { Inter } from '@next/font/google'
const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <div className='bg-slate-700 w-full h-screen flex justify-center items-center'>
        <div className='bg-red-300 w-[600px] h-80 px-24 flex flex-col justify-start items-center'>
          <div className='w-full h-auto flex flex-row justify-center items-center'>
            <Image
                    src="/next.svg"
                    alt="Next.js Logo"
                    width={180}
                    height={37}
                    priority
                  />
          </div>
          <div className='bg-green-300 pt-6'>
            <h2 className='justify-center'>Log in to your notify account</h2>
            <form className='pt-6'>
                <input
                    placeholder="Enter Email"
                    type="text"></input>
                  <input
                    placeholder="Password"
                    type="text"></input>
                  <br></br>
                  <input type="checkbox" name="remember" value="RememberMe"/>
                  <label htmlFor="remember me"> Remember Me</label>
                  <a href="url">Forgot Password?</a>
            </form>
          </div> 
        </div>
      </div>
    </>
  )
}
