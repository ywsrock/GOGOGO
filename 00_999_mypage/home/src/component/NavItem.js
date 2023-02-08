import React, { useEffect, useState } from 'react'
import bgImg from '../assets/bg.png'
import { RiCloseLine, RiMenu3Line } from 'react-icons/ri';
import './NavBar.css'
import { Link, NavLink } from 'react-router-dom';

const NavLinks = ({ ItemClick }) => (
    <>
        <p>
            <NavLink to={`home`} className={({ isActive, isPending }) =>
                isActive
                    ? "active"
                    : isPending
                        ? "pending"
                        : ""
            } onClick={() => ItemClick(false)} >Home</NavLink>
        </p>
        <p>
            <NavLink to={`about`} onClick={() => ItemClick(false)}>About</NavLink>
        </p>
        <p>
            <NavLink to={`news`} onClick={() => ItemClick(false)}>News</NavLink>
        </p>
        <p>
            <NavLink to={`tool`} onClick={() => ItemClick(false)}>Tools</NavLink>
        </p>
    </>
)

const NavItem = () => {
    const [toggenMenu, setToggenMenu] = useState(false);
    const [ws, setWindowWidth] = useState(window.innerWidth);

    useEffect(() => {
        const handleWindowResize = () => {
            setWindowWidth(window.innerWidth);
            setToggenMenu(false);
        };
        window.addEventListener('resize', handleWindowResize);
        return () => window.removeEventListener('resize', handleWindowResize);
    }, [ws])

    return (
        <div className='container'>
            <div className='links'>
                <div className='logo-links'>
                    <Link to='/'>
                        <img src={bgImg} alt='logo'></img>
                    </Link>
                </div>
                <div className='warpper-links'>
                    <NavLinks />
                </div>
            </div>
            <div className='nav-button'>
                <button type='button'>Nav</button>
            </div>
            <div className='nav-menu'>
                {toggenMenu ?
                    <RiCloseLine color='#000000' size={27} onClick={() => setToggenMenu(false)} />
                    : <RiMenu3Line color='#000000' size={27} onClick={() => setToggenMenu(true)} />
                }
                {
                    toggenMenu && (
                        <div className='warpper-menu'>
                            <div className='warpper-menu-links'>
                                <NavLinks SetToggenMenu={setToggenMenu} />
                                <div className='warpper-menu-button'>
                                    <button type='button'>Nav</button>
                                </div>
                            </div>
                        </div>
                    )
                }
            </div>
        </div>
    )
}

export default NavItem